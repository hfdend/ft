package main

import (
    "github.com/hfdend/gohome/site"
    "github.com/hfdend/ft/ctl/index"
)

func main()  {
    site.NewRouter().Add("/", "index.index")

    site.RegisterCtl(&index.Index{})

    site.Run()
}
