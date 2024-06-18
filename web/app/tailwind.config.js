import daisyui from 'daisyui';

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        '../templates/**/*.gotmpl',
        './src/**/*.{ts,js}'
    ],
    theme: {
        extend: {},
        fontFamily: {
            display: ['Playwrite PL', 'cursive'],
            sans: [
                'Montserrat',
                'system-ui',
                '-apple-system',
                'BlinkMacSystemFont',
                'Segoe UI',
                'Roboto',
                'Oxygen',
                'Ubuntu',
                'Cantarell',
                'Open Sans',
                'Helvetica Neue',
                'sans-serif',
            ]
        },
    },
    plugins: [
        daisyui,
    ],
    daisyui: {
        themes: [
            {
                light: {
                    primary: "#69e114",
                    secondary: "#2e99ea",
                    accent: "#e11451",
                    neutral: "#2e1f1f",
                    info: "#a704f7",
                    success: "#1831f6",
                    warning: "#ffd900",
                    error: "#ffad49",
                    "base-100": "#fdfcfc",
                    "base-200": "#e4e3e3",
                    "base-300": "#cacaca",
                    "base-content": "#1c1b1b"
                },
                dark: {
                    primary: "#73eb1e",
                    secondary: "#157fd1",
                    accent: "#eb1e5b",
                    neutral: "#2e1f1f",
                    info: "#a704f7",
                    success: "#1831f6",
                    warning: "#ffd900",
                    error: "#ffad49",
                    "base-100": "#030202",
                    "base-200": "#353535",
                    "base-300": "#686767",
                    "base-content": "#f1f0f0"
                },
            },
        ],
    },
};
