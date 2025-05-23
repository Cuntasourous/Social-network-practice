import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';
import { definePreset } from '@primeuix/themes';  //add this

import App from './App.vue'
import router from './router'
import Button from "primevue/button"

const app = createApp(App)
app.component('Button', Button);

const MyPreset = definePreset(Aura, { //create a preset from Aura and assign colors
    semantic: {
        primary: {
            50: '#f9fafb',
            100: '#f3f4f6',
            200: '#e5e7eb',
            300: '#d1d5db',
            400: '#9ca3af',
            500: '#6b7280',
            600: '#4b5563',
            700: '#374151',
            800: '#1f2937',
            900: '#111827',
            950: '#030712'
        }
    },
    darkMode: false
});

const SlatePreset = definePreset(Aura, { //create a preset from Aura and assign colors
    semantic: {
        primary: {
            50: '{slate.50}',
            100: '{slate.100}',
            200: '{slate.200}',
            300: '{slate.300}',
            400: '{slate.400}',
            500: '{slate.500}',
            600: '{slate.600}',
            700: '{slate.700}',
            800: '{slate.800}',
            900: '{slate.900}',
            950: '{slate.950}'
        }
    }
});


app.use(createPinia())
app.use(router)
//add this
app.use(PrimeVue, {
    theme: {
        preset: SlatePreset //change to Preset name
    }
});

app.mount('#app')
