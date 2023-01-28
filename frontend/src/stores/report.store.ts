import { defineStore } from 'pinia';
import type { IReport } from '@/models/report';
import { getReports } from '@/api/auth';

export type ReportStoreState = {
  reports: IReport[];
  total: Number;
  loading: Boolean;
};

export const useReportStore = defineStore({
  id: 'reportStore',
  state: (): ReportStoreState => ({
    reports: [],
    total: 0,
    loading: false
  }),
  getters: {
    getReports(): IReport[] {
      return this.reports;
    },
    getTotal(): Number {
      return this.total;
    },
    getLoading(): Boolean {
      return this.loading;
    }
  },
  actions: {
    async fetchReports(search: string, start_date: Date, end_date: Date, orderCol: string, order: string, limit: Number) {
      this.reports = [];
      this.loading = true;
      try {
        const response = await getReports(search, start_date, end_date, orderCol, order, limit);
        this.reports = response.data;
        this.total = response.total;
      } finally {
        this.loading = false;
      }
    }
  }
});
