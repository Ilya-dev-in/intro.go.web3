<script>
import { ref } from 'vue';
/* eslint-disable */
export default {
  props: {
    columns: { type: Array },
    tabs: { type: Array },
  },
  computed: {

  },
};
  
/*import { ref, computed, onMounted  } from 'vue'
const count = ref(1);
const name = computed(() => "122" + count.value);

function onClick() {
  count.value++;
}*/
</script>

<template>
   <!-- eslint-disable -->
  <div>
    <table>
      <thead>
        <tr>
          <th v-for="key in columns"
            @click="sortBy(key)"
            :class="{ active: sortKey == key.name }">
            {{ key.name }}
            <!-- <span class="arrow" :class="sortOrders[key.name] > 0 ? 'asc' : 'dsc'"> -->
            <!-- </span> -->
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="entry in data">
          <td v-for="key in columns">
            {{entry[key]}}
          </td>
        </tr>
      </tbody>
    </table>
  </div> 
  
</template>

