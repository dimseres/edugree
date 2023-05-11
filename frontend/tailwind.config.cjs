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
                purple: {
                    DEFAULT: '#6E41E2',
                    50: '#E8E0FA',
                    100: '#DACFF8',
                    200: '#BFABF2',
                    300: '#A488ED',
                    400: '#8964E7',
                    500: '#6E41E2',
                    600: '#4F1FCC',
                    700: '#3C189B',
                    800: '#29106A',
                    900: '#17093A',
                    950: '#0D0521'
                },
                green: {
                    DEFAULT: '#27AE60',
                    50: '#A2EAC1',
                    100: '#92E7B6',
                    200: '#70DF9F',
                    300: '#4FD889',
                    400: '#2ECF72',
                    500: '#27AE60',
                    600: '#1D8047',
                    700: '#12522D',
                    800: '#082514',
                    900: '#000000',
                    950: '#000000'
                },
                red: {
                    DEFAULT: '#DD403A',
                    50: '#F8D8D7',
                    100: '#F5C7C5',
                    200: '#EFA5A2',
                    300: '#E98380',
                    400: '#E3625D',
                    500: '#DD403A',
                    600: '#BE2721',
                    700: '#8E1D19',
                    800: '#5E1310',
                    900: '#2F0908',
                    950: '#170504'
                },
                orange: {
                    DEFAULT: '#D34E24',
                    50: '#F3C9BB',
                    100: '#F0BBAA',
                    200: '#EB9F87',
                    300: '#E58364',
                    400: '#DF6741',
                    500: '#D34E24',
                    600: '#A33C1C',
                    700: '#732B14',
                    800: '#43190B',
                    900: '#130703',
                    950: '#000000'
                },
                yellow: {
                    DEFAULT: '#FFBA49',
                    50: '#FFFFFF',
                    100: '#FFF8EC',
                    200: '#FFE8C3',
                    300: '#FFD99B',
                    400: '#FFC972',
                    500: '#FFBA49',
                    600: '#FFA511',
                    700: '#D88600',
                    800: '#A06300',
                    900: '#684000',
                    950: '#4C2F00'
                }
            },
        },
        fontFamily: {
            sans: ['"Roboto"', 'sans-serif'],
            montserrat: ['"Montserrat"', '"Roboto"', 'sans-serif'],
        },
    },
    plugins: [],
    purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
}
