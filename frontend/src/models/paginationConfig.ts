import type { Ref } from 'vue';

export interface PaginationConfig<T> {
  rowsPerPage?: Ref<number>;
  totalItems?: Ref<number>;
  currentPage: Ref<number>;
}
