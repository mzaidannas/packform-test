<template>
  <section>
    <div class="container mx-auto">
      <div class="flex flex-row pt-8">
        <input type="text" v-model="search" placeholder="Search order, customer or company name" label="Search"
          class="w-full text-white" />
      </div>
      <div class="flex flex-row pt-8">
        <Datepicker v-model="date" range multi-calendars :dark="true" />
      </div>
      <div class="flex flex-row pt-8">
        <span class="text-white">Total amount: ${{ totalAmount }}</span>
      </div>
      <div class="overflow-x-auto sm:-mx-6 lg:-mx-8 pt-6">
        <div class="py-2 align-middle inline-block min-w-full sm:px-3 lg:px-8">
          <div class="shadow overflow-hidden sm:rounded-lg">
            <table class="min-w-full divide-y divide-gray-700 font-semibold text-xs table-fixed">
              <thead class="bg-gray-900">
                <tr>
                  <th scope="col"
                    class="px-3 py-3 text-left text-xs text-gray-300 bg-transparent uppercase tracking-wider whitespace-no-wrap"
                    v-for="(column, id) in columns" :key="id">
                    <span v-if="tableConfig.sortedColumn === id">
                      <svg v-if="tableConfig.sortOrder === 'ASC'" class="w-4 h-4 flex-shrink-0 mr-1 inline"
                        fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                          d="M14.707 12.707a1 1 0 01-1.414 0L10 9.414l-3.293 3.293a1 1 0 01-1.414-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 010 1.414z"
                          clip-rule="evenodd"></path>
                      </svg>
                      <svg v-else class="w-4 h-4 flex-shrink-0 mr-1 inline" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd"
                          d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                          clip-rule="evenodd"></path>
                      </svg>
                    </span>
                    <span class="cursor-pointer" @click="sort(id)">{{ column.name }}</span>
                  </th>
                </tr>
              </thead>
              <tbody class="text-white">
                <tr v-for="(report, id) in reports" :key="id" :class="id % 2 == 0 ? 'bg-gray-800' : 'bg-gray-900'">
                  <td class="text-left px-3 py-3 break-all">
                    {{ report.order_name }}
                  </td>
                  <td class="text-left px-3 py-3 break-all w-2/12">
                    {{ report.customer_company }}
                  </td>
                  <td class="text-left px-3 py-3">
                    {{ report.customer_name }}
                  </td>
                  <td class="px-3 whitespace-no-wrap">
                    {{ report.order_date }}
                  </td>
                  <td class="px-3 whitespace-no-wrap w-1/12">{{ report.delivered_amount || '-' }}</td>
                  <td class="px-3 whitespace-no-wrap w-1/12">${{ report.total_amount }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div class="flex flex-row items-center">
        <span class="text-bold text-white">Total {{ totalPages }}</span>
        <select v-model="limit" class="ml-6 text-white">
          <option selected :value="5">5/page</option>
          <option :value="25">25/page</option>
          <option :value="50">50/page</option>
          <option :value="100">100/page</option>
        </select>
        <vue-awesome-paginate :total-items="total" :items-per-page="tableConfig.limit" :max-pages-shown="3"
          v-model="tableConfig.currentPage" :on-click="onClickHandler"
          pagination-container-class="ml-28"
          paginate-buttons-class="py-1 px-2 cursor-pointer rounded text-white"
          active-page-class="text-blue-600 border-2 border-blue-600 rounded">
          <template #prev-button>
            <div class="px-2 py-1 rounded"
              :class="tableConfig.currentPage > 1 ? ' cursor-pointer text-white bg-blue-600' : 'cursor-default text-gray-500 bg-gray-700'">
              Previous
            </div>
          </template>
          <template #next-button>
            <div
              :class="tableConfig.currentPage < totalPages ? ' cursor-pointer text-white bg-blue-600' : 'cursor-default text-gray-500 bg-gray-700'"
              class="ml-2 px-2 py-1 rounded">
              Next
            </div>
          </template>
        </vue-awesome-paginate>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { useReportStore } from '@/stores/report.store';
import { ref, onMounted, reactive, computed, watch } from 'vue';
import { snakeToPascal } from '@/helpers/formstStrings';

import Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

const { reports, total, loading } = storeToRefs(useReportStore());
const { fetchReports } = useReportStore();

const onClickHandler = async (page: number) => {
  tableConfig.currentPage = page;
  await fetchReports(
    tableConfig.search,
    tableConfig.start_date,
    tableConfig.end_date,
    tableConfig.sortedColumn,
    tableConfig.sortOrder,
    tableConfig.currentPage,
    tableConfig.limit
  );
};

const columns = ref({
  order_name: { name: 'Order_name', align: 'text-left' },
  customer_company: { name: 'Customer Company', align: 'text-left' },
  customer_name: { name: 'Customer Name', align: 'text-left' },
  order_date: { name: 'Order Date', align: 'text-left' },
  delivered_amount: { name: 'Delivered Amount', align: 'text-left' },
  total_amount: { name: 'Total Amount', align: 'text-left' }
});
const date = ref();

const tableConfig = reactive({
  currentPage: 1,
  search: '',
  start_date: new Date(),
  end_date: new Date(),
  sortedColumn: 'order_date',
  sortOrder: 'DESC',
  limit: 5
});

const sortDirs = ['DESC', 'ASC'];

async function sort(column: string) {
  if (tableConfig.sortedColumn == column) {
    tableConfig.sortOrder = sortDirs[(sortDirs.findIndex(col => col == tableConfig.sortOrder) + 1) % sortDirs.length];
  } else {
    tableConfig.sortedColumn = column;
    tableConfig.sortOrder = sortDirs[0];
  }

  // Backend sorting enforces page reset
  onClickHandler(1);
}

onMounted(async () => {
  date.value = [tableConfig.start_date, tableConfig.end_date];
  onClickHandler(1);
});

const totalAmount = computed(() => {
  return reports.value.reduce((sum, report) => sum + report.total_amount.valueOf(), 0);
});

const totalPages = computed(() => {
  return Math.ceil(total.value.valueOf() / tableConfig.limit);
});

const search = ref(tableConfig.search);
watch(search, oldVal => {
  if (timer) {
    clearTimeout(timer);
  }

  var timer = setTimeout(() => {
    tableConfig.search = oldVal;
    // Backend searching enforces page reset
    onClickHandler(1);
  }, 300);
});

watch(date, (oldVal: Date[]) => {
  tableConfig.start_date = oldVal[0];
  tableConfig.end_date = oldVal[1];
  onClickHandler(1);
});

const limit = ref(tableConfig.limit);
watch(limit, (oldVal: Number) => {
  tableConfig.limit = oldVal.valueOf();
  onClickHandler(1);
})
</script>
