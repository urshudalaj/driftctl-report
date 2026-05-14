package renderer

// DarkStyles holds the CSS rules injected when dark mode is active.
type DarkStyles struct {
	Enabled       bool
	Toggle        bool
	DefaultMode   string
	CSS           string
	ToggleScript  string
}

// buildDarkStyles constructs a DarkStyles value from the current Options.
// It returns an empty DarkStyles when dark mode is disabled so that template
// rendering can gate on the Enabled field without extra nil checks.
func buildDarkStyles(o Options) DarkStyles {
	if !o.DarkMode && !o.DarkModeToggle {
		return DarkStyles{}
	}

	defaultMode := o.DarkModeDefault
	if defaultMode == "" {
		defaultMode = "light"
	}

	css := `
:root[data-theme="dark"] {
  --bg: #1a1a2e;
  --surface: #16213e;
  --text: #e0e0e0;
  --muted: #a0a0b0;
  --border: #2a2a4a;
  --accent: #0f3460;
}
`

	script := ""
	if o.DarkModeToggle {
		script = `(function(){
  var s=localStorage.getItem('theme')||'` + defaultMode + `';
  document.documentElement.setAttribute('data-theme',s);
  document.addEventListener('DOMContentLoaded',function(){
    var btn=document.getElementById('theme-toggle');
    if(!btn)return;
    btn.addEventListener('click',function(){
      var cur=document.documentElement.getAttribute('data-theme');
      var next=cur==='dark'?'light':'dark';
      document.documentElement.setAttribute('data-theme',next);
      localStorage.setItem('theme',next);
    });
  });
})()`
	}

	return DarkStyles{
		Enabled:      o.DarkMode || o.DarkModeToggle,
		Toggle:       o.DarkModeToggle,
		DefaultMode:  defaultMode,
		CSS:          css,
		ToggleScript: script,
	}
}
