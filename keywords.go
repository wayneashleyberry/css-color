package parser

// All returns a map of all css color keywords and their associated hex value.
func keywords() map[string]string {
	keywords := make(map[string]string)

	keywords["black"] = "#000000"

	keywords["silver"] = "#c0c0c0"

	keywords["gray"] = "#808080"

	keywords["white"] = "#ffffff"

	keywords["maroon"] = "#800000"

	keywords["red"] = "#ff0000"

	keywords["purple"] = "#800080"

	keywords["green"] = "#008000"

	keywords["lime"] = "#00ff00"

	keywords["olive"] = "#808000"

	keywords["yellow"] = "#ffff00"

	keywords["navy"] = "#000080"

	keywords["blue"] = "#0000ff"

	keywords["teal"] = "#008080"

	keywords["aqua"] = "#00ffff"

	keywords["orange"] = "#ffa500"

	keywords["aliceblue"] = "#f0f8ff"

	keywords["antiquewhite"] = "#faebd7"

	keywords["aquamarine"] = "#7fffd4"

	keywords["azure"] = "#f0ffff"

	keywords["beige"] = "#f5f5dc"

	keywords["bisque"] = "#ffe4c4"

	keywords["blanchedalmond"] = "#ffebcd"

	keywords["blueviolet"] = "#8a2be2"

	keywords["brown"] = "#a52a2a"

	keywords["burlywood"] = "#deb887"

	keywords["cadetblue"] = "#5f9ea0"

	keywords["chartreuse"] = "#7fff00"

	keywords["chocolate"] = "#d2691e"

	keywords["coral"] = "#ff7f50"

	keywords["cornflowerblue"] = "#6495ed"

	keywords["cornsilk"] = "#fff8dc"

	keywords["crimson"] = "#dc143c"

	keywords["cyan"] = "#00ffff"

	keywords["darkblue"] = "#00008b"

	keywords["darkcyan"] = "#008b8b"

	keywords["darkgoldenrod"] = "#b8860b"

	keywords["darkgray"] = "#a9a9a9"

	keywords["darkgreen"] = "#006400"

	keywords["darkgrey"] = "#a9a9a9"

	keywords["darkkhaki"] = "#bdb76b"

	keywords["darkmagenta"] = "#8b008b"

	keywords["darkolivegreen"] = "#556b2f"

	keywords["darkorange"] = "#ff8c00"

	keywords["darkorchid"] = "#9932cc"

	keywords["darkred"] = "#8b0000"

	keywords["darksalmon"] = "#e9967a"

	keywords["darkseagreen"] = "#8fbc8f"

	keywords["darkslateblue"] = "#483d8b"

	keywords["darkslategray"] = "#2f4f4f"

	keywords["darkslategrey"] = "#2f4f4f"

	keywords["darkturquoise"] = "#00ced1"

	keywords["darkviolet"] = "#9400d3"

	keywords["deeppink"] = "#ff1493"

	keywords["deepskyblue"] = "#00bfff"

	keywords["dimgray"] = "#696969"

	keywords["dimgrey"] = "#696969"

	keywords["dodgerblue"] = "#1e90ff"

	keywords["firebrick"] = "#b22222"

	keywords["floralwhite"] = "#fffaf0"

	keywords["forestgreen"] = "#228b22"

	keywords["gainsboro"] = "#dcdcdc"

	keywords["ghostwhite"] = "#f8f8ff"

	keywords["gold"] = "#ffd700"

	keywords["goldenrod"] = "#daa520"

	keywords["greenyellow"] = "#adff2f"

	keywords["grey"] = "#808080"

	keywords["honeydew"] = "#f0fff0"

	keywords["hotpink"] = "#ff69b4"

	keywords["indianred"] = "#cd5c5c"

	keywords["indigo"] = "#4b0082"

	keywords["ivory"] = "#fffff0"

	keywords["khaki"] = "#f0e68c"

	keywords["lavender"] = "#e6e6fa"

	keywords["lavenderblush"] = "#fff0f5"

	keywords["lawngreen"] = "#7cfc00"

	keywords["lemonchiffon"] = "#fffacd"

	keywords["lightblue"] = "#add8e6"

	keywords["lightcoral"] = "#f08080"

	keywords["lightcyan"] = "#e0ffff"

	keywords["lightgoldenrodyellow"] = "#fafad2"

	keywords["lightgray"] = "#d3d3d3"

	keywords["lightgreen"] = "#90ee90"

	keywords["lightgrey"] = "#d3d3d3"

	keywords["lightpink"] = "#ffb6c1"

	keywords["lightsalmon"] = "#ffa07a"

	keywords["lightseagreen"] = "#20b2aa"

	keywords["lightskyblue"] = "#87cefa"

	keywords["lightslategray"] = "#778899"

	keywords["lightslategrey"] = "#778899"

	keywords["lightsteelblue"] = "#b0c4de"

	keywords["lightyellow"] = "#ffffe0"

	keywords["limegreen"] = "#32cd32"

	keywords["linen"] = "#faf0e6"

	keywords["magenta"] = "#ff00ff"

	keywords["fuchsia"] = "#ff00ff"

	keywords["mediumaquamarine"] = "#66cdaa"

	keywords["mediumblue"] = "#0000cd"

	keywords["mediumorchid"] = "#ba55d3"

	keywords["mediumpurple"] = "#9370db"

	keywords["mediumseagreen"] = "#3cb371"

	keywords["mediumslateblue"] = "#7b68ee"

	keywords["mediumspringgreen"] = "#00fa9a"

	keywords["mediumturquoise"] = "#48d1cc"

	keywords["mediumvioletred"] = "#c71585"

	keywords["midnightblue"] = "#191970"

	keywords["mintcream"] = "#f5fffa"

	keywords["mistyrose"] = "#ffe4e1"

	keywords["moccasin"] = "#ffe4b5"

	keywords["navajowhite"] = "#ffdead"

	keywords["oldlace"] = "#fdf5e6"

	keywords["olivedrab"] = "#6b8e23"

	keywords["orangered"] = "#ff4500"

	keywords["orchid"] = "#da70d6"

	keywords["palegoldenrod"] = "#eee8aa"

	keywords["palegreen"] = "#98fb98"

	keywords["paleturquoise"] = "#afeeee"

	keywords["palevioletred"] = "#db7093"

	keywords["papayawhip"] = "#ffefd5"

	keywords["peachpuff"] = "#ffdab9"

	keywords["peru"] = "#cd853f"

	keywords["pink"] = "#ffc0cb"

	keywords["plum"] = "#dda0dd"

	keywords["powderblue"] = "#b0e0e6"

	keywords["rosybrown"] = "#bc8f8f"

	keywords["royalblue"] = "#4169e1"

	keywords["saddlebrown"] = "#8b4513"

	keywords["salmon"] = "#fa8072"

	keywords["sandybrown"] = "#f4a460"

	keywords["seagreen"] = "#2e8b57"

	keywords["seashell"] = "#fff5ee"

	keywords["sienna"] = "#a0522d"

	keywords["skyblue"] = "#87ceeb"

	keywords["slateblue"] = "#6a5acd"

	keywords["slategray"] = "#708090"

	keywords["slategrey"] = "#708090"

	keywords["snow"] = "#fffafa"

	keywords["springgreen"] = "#00ff7f"

	keywords["steelblue"] = "#4682b4"

	keywords["tan"] = "#d2b48c"

	keywords["thistle"] = "#d8bfd8"

	keywords["tomato"] = "#ff6347"

	keywords["turquoise"] = "#40e0d0"

	keywords["violet"] = "#ee82ee"

	keywords["wheat"] = "#f5deb3"

	keywords["whitesmoke"] = "#f5f5f5"

	keywords["yellowgreen"] = "#9acd32"

	keywords["rebeccapurple"] = "#663399"

	return keywords
}
