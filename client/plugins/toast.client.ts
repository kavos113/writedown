import { useToast } from "vue-toast-notification";
import 'vue-toast-notification/dist/theme-default.css';

export default defineNuxtPlugin(() => {
    return {
        provide: {
            toast: useToast({
                position: 'top-right',
                duration: 5000
            })
        }
    };
});