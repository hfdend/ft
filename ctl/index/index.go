package index

import "github.com/hfdend/ft/ctl"

type Index struct {
    ctl.Base
}

func (this *Index) ActionIndex() {

    this.Template.Display()
}
