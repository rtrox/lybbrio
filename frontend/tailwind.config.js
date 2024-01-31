/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    fontFamily: {
      serif: ["Merriweather", "serif"],
      sans: ["Lato", "sans-serif"],
    },
    extend: {
      colors: {
        light: {
          100: "#fcfdfd",
          200: "#f9fbfb",
          300: "#f7f9f8",
          400: "#f4f7f6",
          500: "#f1f5f4", // default
          600: "#cfcfd0", // common grey
          700: "#919392",
          800: "#606262",
          900: "#303131",
        },
        dark: {
          100: "#d4d4d5",
          200: "#a9a9ab",
          300: "#7f7d82",
          400: "#545258",
          500: "#29272e", // default
          600: "#211f25",
          700: "#19171c",
          800: "#111012",
          900: "#080809",
        },
        accent: {
          100: "#fcead2",
          200: "#f9d5a5",
          300: "#f5bf78",
          400: "#f2aa4b",
          500: "#ef951e",
          600: "#bf7718",
          700: "#8f5912",
          800: "#603c0c",
          900: "#301e06",
        },
        primary: {
          100: "#d3e2e2",
          200: "#a6c5c4",
          300: "#7aa9a7",
          400: "#4d8c89",
          500: "#216f6c",
          600: "#1a5956",
          700: "#144341",
          800: "#0d2c2b",
          900: "#071616",
        },
        "light-accent": "#9cc1a2",
        "muted-light-accent": "#a7c0ba",
        success: "#3f9c58",
        warning: "#bc8c20",
        danger: "#f44336",
        default: "#999999",
      },
    },
  },
  plugins: [],
};
