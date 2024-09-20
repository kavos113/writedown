// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2024-04-03',
    devtools: { enabled: true },
    vite: {
        server: {
            proxy: {
                '/api': {
                    target: 'http://localhost:8080',
                    secure: false,
                    changeOrigin: true,
                    rewrite: (path: string) => path.replace(/^\/api/, '')
                }
            }
        }
    },
    runtimeConfig: {
        public: {
            apiBaseURL: process.env.API_BASE_URL || 'http://localhost:8080'
        }
    }
});
