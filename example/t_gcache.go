package main

type gCache struct {
	Title    string
	Duration int
}

//func tGCache1() error {
//	a := gCache{
//		Title:    "test",
//		Duration: 10,
//	}
//	if err := gcache.Set("a", a, 0); err != nil {
//		return err
//	}
//	if v, err := gcache.Get("a"); err != nil {
//		return err
//	} else {
//		b := v.Val().(gCache)
//		xlog.Debug(b)
//		xlog.Debug(b.Title)
//	}
//	return nil
//}
//
//func tGCache2() error {
//	a := &gCache{
//		Title:    "test",
//		Duration: 10,
//	}
//	if err := gcache.Set("a", a, 0); err != nil {
//		return err
//	}
//	if v, err := gcache.Get("a"); err != nil {
//		return err
//	} else {
//		b := v.Val().(*gCache)
//		xlog.Debug(b)
//		xlog.Debug(b.Title)
//	}
//	return nil
//}
//
//func main() {
//	if err := tGCache1(); err != nil {
//		xlog.Error(err)
//	}
//
//	if err := tGCache2(); err != nil {
//		xlog.Error(err)
//	}
//
//}
