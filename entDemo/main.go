package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"v1/ent"
	"v1/ent/car"
	"v1/ent/group"
	"v1/ent/user"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=ent password=gg123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// CreateUser(context.Background(), client)
	// QueryUser(context.Background(), client)
	// CreateCars(context.Background(), client)
	// a8m, _ := QueryUserById(context.Background(), client, 7)
	// a8m := new(ent.User)
	// QueryCars(context.Background(), a8m)
	// QueryCarUsers(context.Background(), a8m)

	// fmt.Println(CreateGraph(context.Background(), client))
	// fmt.Println(QueryGithub(context.Background(), client))
	// fmt.Println(QueryArielCars(context.Background(), client))
	fmt.Println(QueryGroupWithUsers(context.Background(), client))
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %w", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// Get "Ariel" from previous steps.
	a8m := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name("graph"),
		).
		FirstX(ctx)
	cars, err := a8m. // Get the groups, that a8m is connected to:
				QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   //
				Where(         //
			car.Not( //  Get Neta and Ariel cars, but filter out
				car.Model("Mazda"), //  those who named "Mazda"
			), //
		). //
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	fmt.Println("QUERY")
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(23).
		SetName("graph").
		Save(ctx)
	if err != nil {
		return errors.New("a8m error")
	}
	neta, err := client.User.
		Create().
		SetAge(23).
		SetName("neta").
		Save(ctx)
	if err != nil {
		return errors.New("neta error")
	}
	err = client.Car.Create().SetModel("Tesla").SetRegisterAt(time.Now()).
		SetOwner(a8m).Exec(ctx)
	if err != nil {
		return errors.New("a8m setowner error")
	}
	err = client.Car.Create().SetModel("Mazda").SetRegisterAt(time.Now()).
		SetOwner(a8m).Exec(ctx)
	if err != nil {
		return errors.New("neta setowner error")
	}
	err = client.Car.Create().SetModel("Ford").SetRegisterAt(time.Now()).
		SetOwner(neta).Exec(ctx)
	if err != nil {
		return errors.New("neta setowner error")
	}
	err = client.Group.Create().SetName("GitLab").AddUsers(a8m, neta).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return errors.New("set group error")
	}
	err = client.Group.Create().SetName("GitHub").AddUsers(a8m).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return errors.New("set group error")
	}
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// Query the inverse edge.
	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}
	return nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	// user := new(ent.User)
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars:", cars)
	// ford, err := a8m.QueryCars().Where(car.Model("Ford")).Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed querying user cars: %w", err)
	// }
	// log.Println(ford)
	return nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisterAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, errors.New("create car eror")
	}
	log.Println("create cars :", tesla)
	ford, err := client.Car.
		Create().
		SetModel("Car").
		SetRegisterAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, errors.New("create car eror")
	}
	log.Println("create cars :", ford)
	a8m, err := client.User.
		Create().
		SetAge(25).
		SetName("wws").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, errors.New("create car eror")
	}
	log.Println("create users :", a8m)
	return a8m, nil
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUserById(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
