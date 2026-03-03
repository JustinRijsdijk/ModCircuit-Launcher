// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: {
    enabled: true,

    timeline: {
      enabled: true,
    },
  },
  components: true,
  ssr: false,
  css: ['~/assets/css/main.css'],
  nitro: {
    prerender: {
      routes: ["/"],
    },
  },
  modules: ["@nuxtjs/tailwindcss", "@nuxtjs/color-mode", "@nuxt/icon", "@pinia/nuxt"],
  tailwindcss: {
    configPath: 'tailwind.config.ts',
  },
});