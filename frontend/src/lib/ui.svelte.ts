// src/lib/ui.svelte.ts
class UIState {
    isSidebarExpanded = $state(false);
    // Ambil initial value dari localStorage jika ada, jika tidak default ke 'light'
    theme = $state('light'); 

    constructor() {
        // Logika untuk sinkronisasi tema dengan atribut data-theme DaisyUI
        $effect.root(() => {
            $effect(() => {
                if (typeof document !== 'undefined') {
                    document.documentElement.setAttribute('data-theme', this.theme);
                    localStorage.setItem('path-theme', this.theme);
                }
            });
        });
    }

    toggleTheme() {
        this.theme = this.theme === 'light' ? 'dark' : 'light';
    }
}

export const ui = new UIState();