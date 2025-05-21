<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Reactive data properties
const aboutData = ref(null)
const isLoading = ref(true)
const error = ref(null)

// Function to fetch data from Go backend
const fetchAboutData = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/about')
    aboutData.value = response.data
  } catch (err) {
    error.value = 'Failed to load about information'
    console.error('Error:', err)
  } finally {
    isLoading.value = false
  }
}

// Fetch data when component mounts
onMounted(() => {
  fetchAboutData()
})
</script>

<template>
  <div class="about-page">
    <!-- Loading state -->
    <div v-if="isLoading" class="loading">Loading...</div>
    
    <!-- Error state -->
    <div v-if="error" class="error">{{ error }}</div>
    
    <!-- Success state -->
    <div v-if="aboutData" class="about-content">
      <h1>{{ aboutData.message }}</h1>
      <div class="details">
        <p><strong>Version:</strong> {{ aboutData.version }}</p>
        <p><strong>Status:</strong> {{ aboutData.status }}</p>
        <p><strong>Last updated:</strong> 
          {{ new Date(aboutData.timestamp * 1000).toLocaleString() }}
        </p>
      </div>
    </div>
  </div>
</template>