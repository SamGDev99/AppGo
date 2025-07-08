<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import i18next from 'i18next';

i18next.init({
  lng: 'es',
  debug: true,
  resources: {
    es: {
      translation: {
        "Ticker": "Código Acción",
        "Company": "Empresa",
        "Brokerage": "Corredor",
        "Action": "Acción",
        "Rating from": "Calificación inicial",
        "Rating to": "Calificación actual",
        "Target from": "Valor inicial",
        "Target to": "Valor actual"
      }
    }
  }
});

interface Data {
  ID: number;
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  rating_from: string;
  rating_to: string;
  target_from: number;
  target_to: number;
}
type SortableKey = "target_from" | "target_to";
const items = ref<Data[]>([])
const search = ref('')
const currentPage = ref(1)
const itemsPerPage = ref(10)
const loading = ref(true)
let error = ''
const sortKey = ref<'target_from' | 'target_to' | ''>('')
const sortAsc = ref(true)

onMounted(async () => {
  try {
    const res = await fetch('http://localhost:9000/api/datos')
    if (!res.ok) throw new Error('Error al cargar los datos')
    items.value = await res.json()
  } catch (err) {
    if (err instanceof Error) {
      error = err.message
    } else {
      error = 'Error desconocido: ' + err
    }
  } finally {
    loading.value = false
  }
})

const filteredItems = computed(() => {
    let result = items.value

    if(search.value){
        result = result.filter((item) =>
        [item.ticker, item.company, item.brokerage].some(field =>
        field.toLowerCase().includes(search.value.toLowerCase())
        ))
    }

    if (sortKey.value) {
        if (["target_from", "target_to"].includes(sortKey.value)) {
            const key = sortKey.value as SortableKey;
            result = [...result].sort((a, b) => {
            const aVal = a[key];
            const bVal = b[key];
            return sortAsc.value ? aVal - bVal : bVal - aVal;
            });
        }
    }

    return result
}
  

)

const toggleSort = (key: 'target_from' | 'target_to') => {
  if (sortKey.value === key) {
    sortAsc.value = !sortAsc.value
  } else {
    sortKey.value = key
    sortAsc.value = true
  }
}

const totalPages = computed(() =>
  Math.ceil(filteredItems.value.length / itemsPerPage.value)
)

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  return filteredItems.value.slice(start, start + itemsPerPage.value)
})

function goToPage(page: number) {
  if (page >= 1 && page <= totalPages.value) currentPage.value = page
}
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-4 text-gray-800">Listado de Acciones</h1>

    <div class="mb-4 flex items-center justify-between">
      <input
        v-model="search"
        type="text"
        placeholder="Buscar por ticker, empresa o brokerage..."
        class="w-full md:w-1/3 px-4 py-2 border rounded shadow-sm focus:outline-none focus:ring"
      />
      <div class="ml-4">
        <label class="text-sm font-medium mr-2">Por página:</label>
        <select v-model="itemsPerPage" class="border px-2 py-1 rounded">
          <option v-for="n in [10, 25, 50, 100]" :key="n" :value="n">{{ n }}</option>
        </select>
      </div>
    </div>

    <div v-if="loading" class="text-center text-gray-600">Cargando...</div>
    <div v-else-if="error" class="text-red-500">Error: {{ error }}</div>
    <div v-else>
      <table class="min-w-full divide-y divide-gray-200 rounded shadow text-sm">
        <thead class="bg-gray-100 sticky top-0 z-10">
          <tr class="text-left text-gray-700">
            <th class="px-3 py-2">{{i18next.t("Ticker")}}</th>
            <th class="px-3 py-2">{{i18next.t("Company")}}</th>
            <th class="px-3 py-2">{{i18next.t("Brokerage")}}</th>
            <th class="px-3 py-2">{{ i18next.t("Action") }}</th>
            <th class="px-3 py-2">{{ i18next.t("Rating from") }}</th>
            <th class="px-3 py-2">{{ i18next.t("Rating to") }}</th>
            <th class="px-3 py-2 cursor-pointer" @click="toggleSort('target_from')">
                <span v-if="sortKey === 'target_from'">{{ sortAsc ? '↑' : '↓' }}</span>
                {{ i18next.t("Target from") }}
            </th>
            <th class="px-3 py-2 cursor-pointer" @click="toggleSort('target_to')">
                <span v-if="sortKey === 'target_to'">{{ sortAsc ? '↑' : '↓' }}</span>
                {{ i18next.t("Target to") }}
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          <tr
            v-for="item in paginatedItems"
            :key="item.ID"
            class="hover:bg-gray-50"
          >
            <td class="px-3 py-2">{{ item.ticker }}</td>
            <td class="px-3 py-2">{{ item.company }}</td>
            <td class="px-3 py-2">{{ item.brokerage }}</td>
            <td class="px-3 py-2">{{ item.action }}</td>
            <td class="px-3 py-2">{{ item.rating_from }}</td>
            <td class="px-3 py-2">{{ item.rating_to }}</td>
            <td class="px-3 py-2">{{ item.target_from }}</td>
            <td class="px-3 py-2">{{ item.target_to }}</td>
          </tr>
        </tbody>
      </table>

      <div class="mt-4 flex items-center justify-between">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="px-3 py-1 border rounded disabled:opacity-50"
        >
          Anterior
        </button>
        <span class="text-sm">
          Página {{ currentPage }} de {{ totalPages }}
        </span>
        <button
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="px-3 py-1 border rounded disabled:opacity-50"
        >
          Siguiente
        </button>
      </div>
    </div>
  </div>
</template>
    