<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Tabla from './components/Tabla.vue';
import Tarjeta from './components/Tarjeta.vue';
import Titulo from './components/Titulo.vue';

interface Accion {
  ID: number
  ticker: string
  company: string
  brokerage: string
  action: string
  rating_from: string
  rating_to: string
  target_from: number
  target_to: number
  score: number
}

const mejorAccion = ref<Accion | null>(null)
const peorAccion = ref<Accion | null>(null)

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:9000/calcularDatos')
    const data = await response.json()

    if (Array.isArray(data) && data.length === 2) {
      mejorAccion.value = data[0]
      peorAccion.value = data[1]
    }
  } catch (error) {
    console.error('Error al cargar los datos:', error)
  }
})

</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-6xl mx-auto px-4 py-8 space-y-8">
      <Titulo />
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Tarjeta
          v-if="mejorAccion"
          label="Mejor Acción"
          :value="`${mejorAccion.ticker} - $${mejorAccion.target_to}`"
          :score="mejorAccion.score"
          :broker="mejorAccion.brokerage"
          :company-name="mejorAccion.company"
          trend="up"
        />
        <Tarjeta
          v-if="peorAccion"
          label="Peor Acción"
          :value="`${peorAccion.ticker} - $${peorAccion.target_to}`"
          :score="peorAccion.score"
          :broker="peorAccion.brokerage"
          :company-name="peorAccion.company"
          trend="down"
        />
      </div>

      <Tabla />
    </div>
  </div>
</template>

<style scoped>

</style>
