<script setup>
import { ref } from "vue";
const items = ref(['Apple', 'Banana', 'Orange']);
</script>

<template>
    <h2 v-for="item in items" :key="item">
      {{ item }}
    </h2>
</template>