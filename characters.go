package baseEmoji

// From ecnoding/base64.  Needs to be copied here because it's unexported. Also,
// padding character is appended.
const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

// UTF-8 byte order mark
var Utf8Header []byte = []byte{0xEF, 0xBB, 0xBF}

// Faces
var u1F601 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x81})  // Grinning
var u1F602 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x82})  // Joy tears
var u1F603 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x83})  // Open smile
var u1F604 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x84})  // ^^ open smile
var u1F605 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x85})  // sweat drop smile
var u1F606 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x86})  // XD
var u1F609 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x89})  // ;)
var u1F60A Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x8A})  // ^^ closed smile
var u1F60B Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x8B})  // :P "savory food" because that makes sense
var u1F60C Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x8C})  // releaved

var u1F60D Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x8D})  // heart smile
var u1F60F Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x8F})  // smirk
var u1F612 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x92})  // >_>
var u1F613 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x93})  // cold sweat frown
var u1F614 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x94})  // pensive
var u1F616 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x96})  // confused / >_<
var u1F618 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x98})  // kiss
var u1F61A Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x9A})  // kiss w/ closed eyes
var u1F61C Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x9C})  // ;P
var u1F61D Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x9D})  // XP

var u1F61E Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x9E})  // disappointed
var u1F620 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA0})  // angry / >:(
var u1F621 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA1})  // pouting
var u1F622 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA2})  // crying / T_T
var u1F623 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA3})  // perservering
var u1F624 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA4})  // triumph
var u1F625 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA5})  // relieved disappointment
var u1F628 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA8})  // fearfull
var u1F629 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xA9})  // weary
var u1F62A Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xAA})  // sleepy

var u1F62B Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xAB})  // "tired" / DX
var u1F62D Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xAD})  // loud crying
var u1F630 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB0})  // open mouth cold sweat
var u1F631 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB1})  // fear scream
var u1F632 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB2})  // astonished / :O
var u1F633 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB3})  // flushed
var u1F635 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB5})  // dizzy
var u1F637 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB7})  // medical mask
var u1F638 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB8})  // cat grinning
var u1F639 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xB9})  // cat tears of joy

var u1F63A Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBA})  // cat open smile
var u1F63B Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBB})  // cat heart smile
var u1F63C Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBC})  // cat wry smile
var u1F63D Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBD})  // cat kissing
var u1F63E Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBE})  // cat pouting
var u1F63F Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0xBF})  // cat crying
// borked
//var u1F640 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x40})  // cat weary
//var u1F645 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x45})  // no good
//var u1F646 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x46})  // OK
//var u1F647 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x47})  // bowing
//
//var u1F648 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x48})  // see-no-evil
//var u1F649 Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x49})  // hear-no-evil
//var u1F64A Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4A})  // speak-no-evil
//var u1F64B Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4B})  // raised hand / o/
//var u1F64C Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4C})  // raised hands / \o/
//var u1F64D Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4D})  // frowning person
//var u1F64E Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4E})  // pouting person
//var u1F64F Emoji = Emoji([]byte{0xF0, 0x9F, 0x98, 0x4F})  // foulded hands person
// "dingbats"
var u2702 Emoji = Emoji([]byte{0xE2, 0x9C, 0x82})  // black scissors
var u2705 Emoji = Emoji([]byte{0xE2, 0x9C, 0x85})  // black scissors
var u2708 Emoji = Emoji([]byte{0xE2, 0x9C, 0x88})  // airplane
var u2709 Emoji = Emoji([]byte{0xE2, 0x9C, 0x89})  // envelope

var u270A Emoji = Emoji([]byte{0xE2, 0x9C, 0x8A})  // raised fist
var u270B Emoji = Emoji([]byte{0xE2, 0x9C, 0x8B})  // raised hand
var u270C Emoji = Emoji([]byte{0xE2, 0x9C, 0x8C})  // victory hand
var u270F Emoji = Emoji([]byte{0xE2, 0x9C, 0x8F})  // pencil
var u2712 Emoji = Emoji([]byte{0xE2, 0x9C, 0x92})  // fountain pen
var u2714 Emoji = Emoji([]byte{0xE2, 0x9C, 0x94})  // heavy check
var u2716 Emoji = Emoji([]byte{0xE2, 0x9C, 0x96})  // heavy X
var u2728 Emoji = Emoji([]byte{0xE2, 0x9C, 0xA8})  // sparkles
var u2733 Emoji = Emoji([]byte{0xE2, 0x9C, 0xB3})  // asterisk
var u2734 Emoji = Emoji([]byte{0xE2, 0x9D, 0xB4})  // eight-point star

var u2744 Emoji = Emoji([]byte{0xE2, 0x9D, 0x84})  // snowflake
var u2747 Emoji = Emoji([]byte{0xE2, 0x9D, 0x87})  // sparkle
var u274C Emoji = Emoji([]byte{0xE2, 0x9D, 0x8C})  // red X
var u274E Emoji = Emoji([]byte{0xE2, 0x9D, 0x8E})  // negative X

// default padding character
var u2753 Emoji = Emoji([]byte{0xE2, 0x9D, 0x93})  // black question mark

