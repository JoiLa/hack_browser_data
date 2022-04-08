package hack_browser_data

import (
	"fmt"

	"github.com/JoiLa/hack_browser_data/core"
	"github.com/JoiLa/hack_browser_data/core/data"
	"github.com/JoiLa/hack_browser_data/log"
)

/**
 * @description 获取浏览器所有Cookies
 * @author lfs <leyoumake@gmail.com>
 * @param browserName 浏览器名称 ` firefox | firefox-beta | firefox-dev | firefox-nightly | firefox-esr | chrome | edge | brave | chrome-beta | chromium | opera | vivaldi `
 */
func GetBrowserAllCookies(browserName string) (cookieSource *data.CookieSource, err error) {
	log.InitLog("error")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	var browser core.Browser
	browser, err = core.PickOneBrowser(browserName)
	if err != nil {
		return nil, err
	}
	err = browser.InitSecretKey()
	if err != nil {
		return nil, err
	}
	// default select all items
	// you can get single item with browser.GetItem(itemName)
	item, err := browser.GetItem("cookie")
	if err != nil {
		return nil, err
	}
	key := browser.GetSecretKey()
	err = item.CopyDB()
	if err != nil {
		return nil, err
	}
	switch browser.(type) {
	case *core.Chromium:
		err := item.ChromeParse(key)
		if err != nil {
			return nil, err
		}
	case *core.Firefox:
		err := item.FirefoxParse()
		if err != nil {
			return nil, err
		}
	}
	err = item.Release()
	if err != nil {
		return nil, err
	}
	getCookieSource := item.Get().(data.CookieSource)
	cookieSource = &getCookieSource
	return cookieSource, nil
}
