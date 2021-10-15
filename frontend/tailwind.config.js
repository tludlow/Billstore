const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')

module.exports = {
    mode: 'jit',
    purge: ['./components/**/*.{js,jsx,ts,tsx}', './pages/**/*.{js,jsx,ts,tsx}'],
    darkMode: 'media',
    theme: {
        extend: {
            fontFamily: {
                sans: ['Inter var', ...defaultTheme.fontFamily.sans],
            },
            colors: {
                brand: '#3356BC',
                'brand-light': '#4f73db',
                dark: 'rgb(31, 32, 35)',
                'dark-light': 'rgb(39, 40, 43)',
                'dark-lightest': 'rgb(46, 48, 51)',
                gray: colors.trueGray,
            },
        },
    },
    plugins: [require('@tailwindcss/forms'), require('@tailwindcss/ui'), require('@tailwindcss/line-clamp')],
}