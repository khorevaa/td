// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// Types returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		0x5cb934fa: "int32#5cb934fa",
		0xb5286e24: "string#b5286e24",
		0xbc799737: "false#bc799737",
		0x997275b5: "true#997275b5",
		0xe937bb82: "bytes#e937bb82",
		0x14feebbc: "error#14feebbc",
		0xd4edbe69: "ok#d4edbe69",
		0xec200d96: "message#ec200d96",
		0xed8bebfe: "sms#ed8bebfe",
		0x85d7fd8b: "responseID#85d7fd8b",
		0xcb0244f2: "responseText#cb0244f2",
		0x7490dcc5: "bigMessage#7490dcc5",
		0xee6324c4: "noMessage#ee6324c4",
		0xcc6136f1: "targetsMessage#cc6136f1",
		0xb03e2ef8: "update#b03e2ef8",
		0x2b4b45c:  "getUpdatesResp#2b4b45c",
		0x947225b5: "fieldsMessage#947225b5",
		0xf990a67d: "bytesMessage#f990a67d",
		0x37b3df65: "textEntityTypeMention#37b3df65",
		0xc2f7a2dd: "textEntityTypeHashtag#c2f7a2dd",
		0x48e4374b: "textEntityTypeCashtag#48e4374b",
		0xbb652bb3: "textEntityTypeBotCommand#bb652bb3",
		0xb1c0d47c: "textEntityTypeUrl#b1c0d47c",
		0x54f81821: "textEntityTypeEmailAddress#54f81821",
		0xbad9aa2a: "textEntityTypePhoneNumber#bad9aa2a",
		0x6513910:  "textEntityTypeBankCardNumber#6513910",
		0xbcc0e1b0: "textEntityTypeBold#bcc0e1b0",
		0xf8f3965d: "textEntityTypeItalic#f8f3965d",
		0x2f39cf92: "textEntityTypeUnderline#2f39cf92",
		0x394fc4fa: "textEntityTypeStrikethrough#394fc4fa",
		0xc5e9c94a: "textEntityTypeCode#c5e9c94a",
		0x62491c8e: "textEntityTypePre#62491c8e",
		0xc7a77aab: "textEntityTypePreCode#c7a77aab",
		0x1a912463: "textEntityTypeTextUrl#1a912463",
		0xd0d2685d: "textEntityTypeMentionName#d0d2685d",
		0x8bab99a8: "textEntity#8bab99a8",
		0xcf89c258: "textEntities#cf89c258",
		0xddbd2c09: "testInt#ddbd2c09",
		0xfe56688c: "testString#fe56688c",
		0xa422c4de: "testBytes#a422c4de",
		0xdf9eb113: "testVectorInt#df9eb113",
		0xf152999b: "testVectorIntObject#f152999b",
		0x5d6f85bc: "testVectorString#5d6f85bc",
		0xe5ecc0d:  "testVectorStringObject#e5ecc0d",
		0xa590fb25: "testVectorBytes#a590fb25",
		0x69e8846c: "testVectorVector#69e8846c",
		0x6643b654: "client_DH_inner_data#6643b654",
		0x18b7a10d: "dcOption#18b7a10d",
		0x330b4067: "config#330b4067",
		0xf8bb4a38: "auth#f8bb4a38",
		0x29bacabb: "authPassword#29bacabb",
		0xf4815592: "user.auth#f4815592",
		0x5981e317: "user.authPassword#5981e317",
		0x28f1114:  "theme#28f1114",
		0xf41eb622: "account.themesNotModified#f41eb622",
		0x7f676421: "account.themes#7f676421",
		0xce73048f: "ping#ce73048f",
		0xf74488a:  "send#f74488a",
		0xdf18e5ca: "sendMultipleSMS#df18e5ca",
		0xfd2f6687: "doAuth#fd2f6687",
		0xd4785939: "echoVector#d4785939",
	}
}
