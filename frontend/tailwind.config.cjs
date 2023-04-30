/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [],
    theme: {
        screens: {
            sm: '480px',
            md: '768px',
            lg: '976px',
            xl: '1440px',
        },
        extend: {
            colors: {
                'blue': '#6E41E2',
                'green': '#27AE60',
                'red': '#DD403A',
                'orange': '#D34E24',
                'yellow': '#FFBA49',
                'gray-dark': '#3E363F',
                'gray': '#605F5E',
                'gray-light': '#D8DBE2',
                'white': '#FFFFFF',
            },
        },
        fontFamily: {
            sans: ['"Roboto"', 'sans-serif'],
            montseratt: ['"Montserrat"', '"Roboto"', 'sans-serif'],
        },
    },
    plugins: [],
    purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
}
