package cli

// https://en.wikipedia.org/wiki/Miscellaneous_Symbols
// https://www.w3schools.com/charsets/ref_html_utf8.asp

// TODO: This will be moved into its own package and called in optionally

// A collection of UTF-8 symbols that work by default in Gnome terminal on
// Debian AND that are specifically useful for UI design.

//
// Checkbox

var checkbox = map[string]string{
	//  ▣
	"empty": "☐",
	"check": "☑",
	"minus": "⊟",
	"plus":  "⊞",
	"times": "⊠",
	//◼
}

// Greek
// α β δ ε θ λ μ π φ ψ Ω

// Logic
//  ¬ ⊨ ⊭ ∀ ∁ ∃ ∄ ∴ ∵ ⊦ ⊬ ⊧ ⊩ ⊮ ⊫ ⊯ ⊪

// Subgroups
// ⊲ ⊳ ⊴ ⊵ ⋪ ⋫ ⋬ ⋭

// Approx
//  ≅≆ ≇  ≈≉ ≊

// Identify
// ≡ ≢

// Equiv
// ≍ ≭ ≣

// Phonetic
//  ā ä ē ĕ ī ĭ ō ŏ ȯ ô û ü

// Arrows
// ← → ↑ ↓
//↔ ↕
//↖ ↗ ↘ ↙
//⤡ ⤢
//↚ ↛ ↮
//⟵ ⟶ ⟷
//⇦ ⇨ ⇧ ⇩
//⬄ ⇳
//⬁ ⬀ ⬂ ⬃
//⬅ ➡  ⬆ ⬇
//⬉ ⬈ ⬊ ⬋
//⬌ ⬍

// ? bowties?
//  ⋈ ⧑ ⧒ ⧓ ⧔ ⧕

// Checkers
// ⛀ ⛁ ⛂ ⛃

// Dice
// ⚀ ⚁ ⚂ ⚃ ⚄ ⚅

// Triangle Arrows
// ▲ ▼ ◀ ▶
//△ ▽ ◁ ▷
//▴ ▾ ◂ ▸
//▵ ▿ ◃ ▹

var warning = map[string]string{
	"skull":       "☠",
	"radioactive": "☢",
	"caution":     "☡",
	"biohazard":   "☣",
	"warning":     "⚠",
}

var star = map[string]string{
	"filled": "★",
	"empty":  "☆",
}

// patterns
// ▤ ▥ ▦ ▧ ▨ ▩

// Could be useful for indicating corners
//  ◰ ◱ ◲ ◳

// Signs
// ⚛ ☠ ☢ ☣ ⚡
//♻ ♼ ♽ ♲ ♾
//♺ ♳ ♴ ♵ ♶ ♷ ♸ ♹
// ♿

var flag = map[string]string{
	"filled": "⚑",
	"empty":  "⚐",
}

var heart = map[string]string{
	"filled": "♡",
	"empty":  "♡",
}

var settings = map[string]string{
	"brightness": "☼",
}

//"hermaphrodite": "⚦",
var sex = map[string]string{
	"female": "♀",
	"male":   "♂",
	"trans":  "⚧", // or ⚦  or  ⚥ or ⚨
	"inter":  "⚩",
}

// Currencies
// ¤
//  $ ¢ € ₠ £ ₨ ₹ ₵ ₡
//₳ ฿ ₣ ₲ ₭ ₥ ₦ ₱ ₽ ₴ ₮ ₩ ¥
//  ₢ ₫ ₯ ₪ ₧ ₰

// tech
// telephone ☎ ☏ ✆ ℡
// keyboard
// tape drive ✇

//var numbers = map[string]map[string]string{
//	⓪ ① ② ③ ④ ⑤ ⑥ ⑦ ⑧ ⑨ ⑩ ⑪ ⑫ ⑬ ⑭ ⑮ ⑯ ⑰ ⑱ ⑲ ⑳
//   ⓿ ❶ ❷ ❸ ❹ ❺ ❻ ❼ ❽ ❾ ❿ ⓫ ⓬ ⓭ ⓮ ⓯ ⓰ ⓱ ⓲ ⓳ ⓴
//   ➊ ➋ ➌ ➍ ➎ ➏ ➐ ➑ ➒ ➓ - sans serif
//}

var chess = map[string]map[string]string{
	"white": map[string]string{
		"king":   "♔",
		"queen":  "♕",
		"rook":   "♖",
		"bishop": "♗",
		"knight": "♘",
		"pawn":   "♙",
	},
	"black": map[string]string{
		"king":   "♚",
		"queen":  "♛",
		"rook":   "♜",
		"bishop": "♝",
		"knight": "♞",
		"pawn":   "♟",
	},
}

//  🂱 🂲 🂳 🂴 🂵 🂶 🂷 🂸 🂹 🂺
//  🂱 🂲 🂳 🂴 🂵 🂶 🂷 🂸 🂹 🂺
// 🂻 🂼 🂽 🂾
// 🂡 🂢 🂣 🂤 🂥 🂦 🂧 🂨 🂩 🂪
// 🂫 🂬 🂭 🂮
// 🃁 🃂 🃃 🃄 🃅 🃆 🃇 🃈 🃉 🃊
// 🃋 🃌 🃍 🃎
// 🃑 🃒 🃓 🃔 🃕 🃖 🃗 🃘 🃙 🃚
// 🃛 🃜 🃝 🃞
// 🂠 🃏 🃟
var cards = map[string]map[string]string{
	"spade": map[string]string{
		"empty":  "♣",
		"filled": "♠",
	},
	"heart": map[string]string{
		"empty":  "♡",
		"filled": "♡",
	},
	"diamond": map[string]string{
		"empty":  "♢",
		"filled": "",
	},
	"club": map[string]string{
		"empty":  "♣",
		"filled": "♣",
	},
}

var face = map[string]map[string]string{
	"empty": map[string]string{
		"frowning": "☹",
		"smile":    "☺",
	},
	"filled": map[string]string{
		"smile": "☻",
	},
}

var equals = map[string]string{
	"tilde": "⋍",
}

// Out of below decide what makes sense, just use ones that would likely see use
// in a UI. Like math makes sense

// Keyboard
// ⌘ ⌥ ⌫ ⌦ ⏏ ⏎ ❖

// Math
// ∅ ∈ ∉ ⊂ ⊃ ⊆ ⊇ ⊄ ⋂ ⋃ ≠ ≤ ≥ ≈ ≡ ⌈ ⌉ ⌊ ⌋ ∑ ∫ × ⊗ ⊕ ⊙ × ∂ √ ≔ ± ² ℵ ∞ ∎ ′ ° ℓ ∡ ∀ ¬ ∧ ∨ ∃ ∵ ∴
//  % ‰ ‱
// ƒ

// Helm key (boat steering wheel, used on french macbook for ctrl)
// ⎈

// Arrows
// ⇧⇪⇪⇫⇬⇮⇯⇭
// ↩↵⏎ ↹ ⇄ ⇤ ⇥ ↤ ↦
// ↖ ↘ ⇤ ⇥ ⤒ ⤓ ⇱ ⇲ ⟵
// ↑ ↓ ← →
// ⇦ ⇨ ⇧ ⇩ ⬅  ➡⬆ ⬇
// △ ▽ ▲ ▼  ◀  ▶▲ ▼ ◁ ▷ △ ▽

// Stacked Lines
// ▤ ☰ 𝌆

// Checks
// ✓ ✔ ✗ ✘

// Numbersign
// №

// Bullets
// • ◦ ‣ ⁃ ◘ ❥ ⁌ ⁍ ☙ ❧

// Editing
// ⁀ ⎁ ⎂ ⎃ �

// White space representation
//tab ↹ ⇄ ⇤ ⇥ ↤ ↦ ◁ ▷
//space
//· ␣ ˽  ␢
//paragraph, section
//¶ § ␤

// Punctuation
// ‼

// Various
// ^ ⌃
// ✲
//  ⎇ ⌥ ✦ ✧   ⌤
// ⎋ ⌫  ⌦ ⎀
// ⌧  ⇞ ⇟   ⎉ ⎊ ⍰  ☾
// ⌂

// Info
// ℹ

// Letter
// ✉

// Cut
// ✂ ✄

// Reload
// ↶ ↷ ⟲ ⟳ ↺ ↻

// Boxes
// ┌ ─ ┬ ┐
//│ │ │ │
//├ ─ ┼ ┤
//└ ─ ┴ ┘
//┏ ┳ ━ ┓
//┣ ╋ ━ ┫
//┃ ┃ ┃ ┃
//┗ ┻ ━ ┛
//╔ ╦ ═ ╗
//╠ ╬ ═ ╣
//║ ║ ║ ║
//╚ ╩ ═ ╝
//╒ ╤ ╕
//╞ ╪ ╡
//╘ ╧ ╛
//╓ ╥ ╖
//╟ ╫ ╢
//╙ ╨ ╜
//┍ ┯ ┑
//┝ ┿ ┥
//┕ ┷ ┙
//╆ ╅
//╄ ╃
//┲ ┱
//┺ ┹
//┢ ╈ ┪
//╊ ╋ ╉
//┡ ╇ ┩
//┎ ┰ ┒
//┠ ╂ ┨
//┖ ┸ ┚
//┟ ╁ ┧
//┞ ╀ ┦
//┮ ┭
//┾ ┽
//┶ ┵
//╌ ╎ ╍ ╏
//┄ ┆ ┊ ┈
//┅ ┇ ┉ ┋
//╭ ╮
//╰ ╯ ╱ ╲ ╳
//╴ ╶ ╵ ╷
//╸ ╺ ╹ ╻
//╼ ╾ ╽ ╿  ┌ ─ ┬ ┐
//│ │ │ │
//├ ─ ┼ ┤
//└ ─ ┴ ┘
//┏ ┳ ━ ┓
//┣ ╋ ━ ┫
//┃ ┃ ┃ ┃
//┗ ┻ ━ ┛
//╔ ╦ ═ ╗
//╠ ╬ ═ ╣
//║ ║ ║ ║
//╚ ╩ ═ ╝
//╒ ╤ ╕
//╞ ╪ ╡
//╘ ╧ ╛
//╓ ╥ ╖
//╟ ╫ ╢
//╙ ╨ ╜
//┍ ┯ ┑
//┝ ┿ ┥
//┕ ┷ ┙
//╆ ╅
//╄ ╃
//┲ ┱
//┺ ┹
//┢ ╈ ┪
//╊ ╋ ╉
//┡ ╇ ┩
//┎ ┰ ┒
//┠ ╂ ┨
//┖ ┸ ┚
//┟ ╁ ┧
//┞ ╀ ┦
//┮ ┭
//┾ ┽
//┶ ┵
//╌ ╎ ╍ ╏
//┄ ┆ ┊ ┈
//┅ ┇ ┉ ┋
//╭ ╮
//╰ ╯ ╱ ╲ ╳
//╴ ╶ ╵ ╷
//╸ ╺ ╹ ╻
//╼ ╾ ╽ ╿

// Domino
//  🁣 🁤 🁥 🁦 🁧 🁨 🁩
// 🁪 🁫 🁬 🁭 🁮 🁯 🁰
// 🁱 🁲 🁳 🁴 🁵 🁶 🁷
// 🁸 🁹 🁺 🁻 🁼 🁽 🁾
// 🁿 🂀 🂁 🂂 🂃 🂄 🂅
// 🂆 🂇 🂈 🂉 🂊 🂋 🂌
// 🂍 🂎 🂏 🂐 🂑 🂒 🂓
// 🁢
// 🀱 🀲 🀳 🀴 🀵 🀶 🀷
// 🀸 🀹 🀺 🀻 🀼 🀽 🀾
// 🀿 🁀 🁁 🁂 🁃 🁄 🁅
// 🁆 🁇 🁈 🁉 🁊 🁋 🁌
// 🁍 🁎 🁏 🁐 🁑 🁒 🁓
// 🁔 🁕 🁖 🁗 🁘 🁙 🁚
// 🁛 🁜 🁝 🁞 🁟 🁠 🁡
// 🀰
