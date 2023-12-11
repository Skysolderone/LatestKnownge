pacakge main

type Post struct {
	UID     string `json:"uid" gorm:"primary"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug" gorm:"uniqueIndex`
 }
 type NewPostMessage struct {
	UID     string `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
 }
 
 type PublishedPostMessage struct {
	Post
 }

 type Publisher struct {
	newPostReader       *kafka.Reader
	publishedPostWriter *kafka.Writer
	db                  *gorm.DB
 }
func main(){

}


func NewPublisher() (*Publisher, func()) {
	// we should use dependency injection here, but the point of the article is something else, therefore, for
	// simplicity, I initiated all dependencies here.
	p := &Publisher{}
	mechanism, err := scram.Mechanism(scram.SHA256, "","")
	if err != nil {
	   log.Fatalln(err)
	}
 
	// setup database
	dsn := "host=localhost user= password= dbname=posts sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 
	if err != nil {
	   log.Fatalln(err)
	}
	p.db = db
	if err := db.AutoMigrate(&Post{}); err != nil {
	   log.Fatalln(err)
	}
 
	// setup kafka
	dialer := &kafka.Dialer{SASLMechanism: mechanism, TLS: &tls.Config{}}
	p.newPostReader = kafka.NewReader(kafka.ReaderConfig{
	   Brokers: []string{""},
	   Topic:   "app.newPosts",
	   GroupID: "service.publisher",
	   Dialer:  dialer,
	})
	p.publishedPostWriter = kafka.NewWriter(kafka.WriterConfig{
	   Brokers: []string{""},
	   Topic:   "app.publishedPosts",
	   Dialer:  dialer,
	})
 
	return p, func() {
	   p.newPostReader.Close()
	   p.publishedPostWriter.Close()
	}
 }

 func (p *Publisher) Run() {
	for {
	   newPost, err := p.newPostReader.FetchMessage(context.Background())
	   if err != nil {
		  if errors.Is(err, io.EOF) {
			 return
		  }
		  log.Fatalln(err)
	   }
 
	   var post NewPostMessage
	   if err := json.Unmarshal(newPost.Value, &post); err != nil { // dead letter queue maybe a better solution
		  log.Printf("decoding new post error: %s\n", err.Error())
		  continue
	   }
 
	   postModel := Post{
		  UID:     post.UID,
		  Title:   post.Title,
		  Content: post.Content,
		  Slug:    slug.Make(post.Title + "-" + time.Now().Format(time.Stamp)),
	   }
	   if err := p.db.Create(&postModel).Error; err != nil {
		  log.Printf("saving new post in database: %s\n", err.Error())
	   }
	   p.newPostReader.CommitMessages(context.Background(), newPost)
 
	   b, _ := json.Marshal(PublishedPostMessage{Post: postModel})
	   p.publishedPostWriter.WriteMessages(context.Background(), kafka.Message{Value: b})
	   log.Printf("the %s post has been saved in the database\n", post.UID)
	}
 }

 type CacheManager struct {
	publishedPostReader *kafka.Reader
	rdb                 *redis.Client
 }
 
 func NewCacheManager() (*CacheManager, func()) {
	// we should use dependency injection here, but the point of the article is something else, therefore, for
	// simplicity, I initiated all dependencies here.
	cm := &CacheManager{}
	mechanism, err := scram.Mechanism(scram.SHA256, "", "")
	if err != nil {
	   log.Fatalln(err)
	}
 
	// setup Redis
	opt, _ := redis.ParseURL("redis://default:PASSWORD@SERVER:PORT")
	cm.rdb = redis.NewClient(opt)
 
	// setup kafka
	dialer := &kafka.Dialer{SASLMechanism: mechanism, TLS: &tls.Config{}}
	cm.publishedPostReader = kafka.NewReader(kafka.ReaderConfig{
	   Brokers: []string{""},
	   Topic:   "app.publishedPosts",
	   GroupID: "service.cacheManager",
	   Dialer:  dialer,
	})
 
	return cm, func() {
	   cm.publishedPostReader.Close()
	   cm.rdb.Close()
	}
 }
 
 func (c *CacheManager) Run() {
	for {
	   publishedPost, err := c.publishedPostReader.FetchMessage(context.Background())
	   if err != nil {
		  if errors.Is(err, io.EOF) {
			 return
		  }
		  log.Fatalln(err)
	   }
 
	   var published PublishedPostMessage
	   if err := json.Unmarshal(publishedPost.Value, &published); err != nil {
		  log.Printf("decoding published post error: %s\n", err.Error())
		  continue
	   }
 
	   b, _ := json.Marshal(published.Post)
	   c.rdb.Set(context.Background(), "post:"+published.Slug, b, 0)
	   c.publishedPostReader.CommitMessages(context.Background(), publishedPost)
	   log.Printf("the %s post has been saved in Redis\n", published.UID)
	}
 }