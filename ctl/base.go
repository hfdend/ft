package ctl

import "github.com/hfdend/gohome/site"

type Base struct {
    site.Ctl
}

func (this *Base) ActionPre() {
    this.Template.AddCommTemplateFile("header.html")
    this.Template.AddCommTemplateFile("footer.html")
    this.Ctl.ActionPre()
}

func init()  {
}