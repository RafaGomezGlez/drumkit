import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

const API_GW = 'https://nnef6rysh7.execute-api.us-east-1.amazonaws.com'

// Updated interface to match your actual API response
export interface LoadData {
  id: number;
  customId: string;
  status: {
    code: {
      key: string;
      value: string;
    };
  };
  customerOrder: Array<{
    id: number;
    customer: {
      id: number;
      name: string;
      parentAccount: {
        name: string;
        type: string;
        id: number;
      };
    };
    deleted: boolean;
  }>;
  carrierOrder: any | null;
  created: string;
  updated: string;
  lastUpdatedOn: string;
  createdDate: string;
}

// Keep the old interface for create operations
export interface Load {
  id?: string
  pickup: {
    name: string
    apptTime: string
    city: string
    state: string
    country: string
  }
  consignee: {
    name: string
    apptTime: string
    city: string
    state: string
    country: string
  }
  status: string
  customer: {
    name: string
    externalTMSId: string
  }
  specifications: {
    minTempFahrenheit: number
    maxTempFahrenheit: number
  }
  totalWeight: number
}

export interface ViewLoadsParams {
  start?: string
  pageSize?: string
}

export interface ViewLoadsResponse extends Array<LoadData> {}

export const drumkitAPI = createApi({
  reducerPath: 'drumkitAPI',
  baseQuery: fetchBaseQuery({
    baseUrl: `${API_GW}/v2`,
    headers: {
      'Content-Type': 'application/json',
    },
  }),
  tagTypes: ['Load'],
  endpoints: (builder) => ({
    viewLoads: builder.query({
      query: (params?: ViewLoadsParams) => {
        const searchParams = new URLSearchParams()
        if (params?.start) {
          searchParams.append('start', params.start)
        }
        if (params?.pageSize) {
          searchParams.append('pageSize', params.pageSize)
        }
        
        const queryString = searchParams.toString()
        
        return {
          url: `view-loads${queryString ? `?${queryString}` : ''}`,
          method: 'GET',
        }
      },
      providesTags: ['Load'],
    }),
    createLoad: builder.mutation({
      query: (loadData: Load) => ({
        url: 'create-load',
        method: 'POST',
        body: loadData,
      }),
      invalidatesTags: ['Load'],
    }),
  }),
})

export const { useViewLoadsQuery, useCreateLoadMutation } = drumkitAPI
