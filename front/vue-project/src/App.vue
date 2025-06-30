<script setup lang="ts">
import Tabla from './components/Tabla.vue';
import Tarjeta from './components/Tarjeta.vue';
import Titulo from './components/Titulo.vue';

const mejorAccion = {
  ticker: 'AAPL',
  target_to: 230.15,
  variacion: 67.81
}

const peorAccion = {
  ticker: 'TSLA',
  target_to: 100.55,
  variacion: -42.24
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-6xl mx-auto px-4 py-8 space-y-8">
      <Titulo />
      
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <Tarjeta
          label="Mejor Acción"
          :value="mejorAccion.ticker + ' - $' + mejorAccion.target_to"
          :percentage="mejorAccion.variacion"
          trend="up"
        />
        <Tarjeta
          label="Peor Acción"
          :value="peorAccion.ticker + ' - $' + peorAccion.target_to"
          :percentage="peorAccion.variacion"
          trend="down"
        />
      </div>

      <Tabla />
    </div>
  </div>
</template>

<style scoped>

</style>
