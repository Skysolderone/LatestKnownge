package main

import (
	"fmt"

	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
)

func main() {
	pkg, _ := apk.OpenFile("testUtrading.apk")
	defer pkg.Close()

	// icon, _ := pkg.Icon(nil)     // returns the icon of APK as image.Image
	pkgName := pkg.PackageName() // returns the package name

	resConfigEN := &androidbinary.ResTableConfig{
		Language: [2]uint8{uint8('e'), uint8('n')},
	}
	appLabel, _ := pkg.Label(resConfigEN) // get app label for en translation

	// fmt.Println(icon)
	fmt.Println(string(pkgName))
	fmt.Println(appLabel)
}
