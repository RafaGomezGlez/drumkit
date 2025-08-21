import React from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Tag } from 'primereact/tag';
import { Card } from 'primereact/card';
import { typography, colors, textStyles } from '../styles/typography';

// Custom styles for minimalist table with improved typography
const tableStyles = `
  .p-datatable-minimalist {
    font-family: ${typography.fonts.primary};
  }

  .p-datatable-minimalist .p-datatable-tbody > tr {
    background: ${colors.background.primary};
    transition: all 0.2s ease;
    border-bottom: 1px solid ${colors.border.light};
  }
  
  .p-datatable-minimalist .p-datatable-tbody > tr:hover {
    background: ${colors.background.secondary} !important;
  }
  
  .p-datatable-minimalist .p-datatable-thead > tr > th {
    background-color: ${colors.background.secondary};
    color: ${colors.text.secondary};
    font-weight: ${typography.weights.medium};
    font-size: ${typography.sizes.xs};
    text-transform: uppercase;
    letter-spacing: ${typography.letterSpacing.wide};
    padding: 1.25rem 1rem;
    border-bottom: 2px solid ${colors.border.light};
    font-family: ${typography.fonts.primary};
  }
  
  .p-datatable-minimalist .p-datatable-tbody > tr > td {
    padding: 1.5rem 1rem;
    border-bottom: 1px solid ${colors.border.light};
    font-size: ${typography.sizes.base};
    color: ${colors.text.primary};
  }
  
  .p-datatable-minimalist .p-paginator {
    background: transparent;
    border: none;
    padding: 1.5rem 0;
    font-family: ${typography.fonts.primary};
  }
  
  .p-datatable-minimalist .p-paginator .p-paginator-pages .p-paginator-page {
    background: transparent;
    border: 1px solid ${colors.border.light};
    color: ${colors.text.secondary};
    margin: 0 4px;
    border-radius: 6px;
    font-size: ${typography.sizes.sm};
    font-weight: ${typography.weights.medium};
  }
  
  .p-datatable-minimalist .p-paginator .p-paginator-pages .p-paginator-page:hover {
    background: ${colors.background.tertiary};
  }
  
  .p-datatable-minimalist .p-paginator .p-paginator-pages .p-paginator-page.p-highlight {
    background: ${colors.brand.primary};
    border-color: ${colors.brand.primary};
    color: white;
  }
  
  .p-datatable-minimalist .p-column-filter-row .p-column-filter-element {
    width: 100%;
  }
  
  .p-datatable-minimalist .p-column-filter-row .p-inputtext {
    border: 1px solid ${colors.border.light};
    border-radius: 8px;
    padding: 0.75rem 1rem;
    font-size: ${typography.sizes.sm};
    font-family: ${typography.fonts.primary};
    background: ${colors.background.primary};
    color: ${colors.text.primary};
    width: 100%;
    transition: all 0.2s ease;
  }
  
  .p-datatable-minimalist .p-column-filter-row .p-inputtext:focus {
    border-color: ${colors.brand.primary};
    box-shadow: 0 0 0 2px rgba(25, 118, 210, 0.1);
    outline: none;
  }
  
  .p-datatable-minimalist .p-column-filter-row .p-inputtext::placeholder {
    color: ${colors.text.tertiary};
    font-size: ${typography.sizes.xs};
  }
  
  .p-datatable-minimalist .p-datatable-thead > tr.p-filter-row > th {
    background: ${colors.background.primary};
    padding: 1rem;
    border-bottom: 1px solid ${colors.border.light};
  }
`;

// Define the interface based on your API response
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

interface TableViewProps {
  data: LoadData[];
  loading?: boolean;
  totalRecords?: number;
  first?: number;
  rows?: number;
  onPage?: (event: any) => void;
}

const TableView: React.FC<TableViewProps> = ({ 
  data = [], 
  loading = false, 
  totalRecords = 0,
  first = 0,
  rows = 25,
  onPage
}) => {
  
  // Inject custom styles
  React.useEffect(() => {
    const styleElement = document.createElement('style');
    styleElement.textContent = tableStyles;
    document.head.appendChild(styleElement);
    
    return () => {
      document.head.removeChild(styleElement);
    };
  }, []);
  
  // Format date to readable format
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  // Status template with color coding
  const statusTemplate = (rowData: LoadData) => {
    const status = rowData.status.code.value;
    const severity = status === 'Covered' ? 'success' : 
                    status === 'Tendered' ? 'warning' : 'info';
    
    return (
      <Tag 
        value={status} 
        severity={severity} 
        rounded 
        style={{ 
          padding: '0.4rem 0.8rem',
          fontSize: typography.sizes.xs,
          fontWeight: typography.weights.medium,
          letterSpacing: typography.letterSpacing.wide,
          fontFamily: typography.fonts.primary
        }} 
      />
    );
  };

  // Customer template
  const customerTemplate = (rowData: LoadData) => {
    const customer = rowData.customerOrder[0]?.customer;
    if (!customer) return '-';
    
    return (
      <div style={{ lineHeight: typography.lineHeights.tight }}>
        <div style={{ 
          fontWeight: typography.weights.medium, 
          color: colors.text.primary,
          marginBottom: '2px',
          fontSize: typography.sizes.base,
          fontFamily: typography.fonts.primary
        }}>
          {customer.name}
        </div>
        <div style={{ 
          fontSize: typography.sizes.xs, 
          color: colors.text.secondary,
          fontFamily: typography.fonts.monospace
        }}>
          ID: {customer.id}
        </div>
      </div>
    );
  };

  // Date template
  const dateTemplate = (rowData: LoadData) => {
    return (
      <div style={{ 
        fontSize: typography.sizes.sm,
        color: colors.text.secondary,
        fontFamily: typography.fonts.primary
      }}>
        {formatDate(rowData.created)}
      </div>
    );
  };

  // Updated date template
  const updatedTemplate = (rowData: LoadData) => {
    return (
      <div style={{ 
        fontSize: typography.sizes.sm,
        color: colors.text.secondary,
        fontFamily: typography.fonts.primary
      }}>
        {formatDate(rowData.lastUpdatedOn)}
      </div>
    );
  };

  // ID template with custom formatting
  const idTemplate = (rowData: LoadData) => {
    return (
      <div style={{ lineHeight: typography.lineHeights.tight }}>
        <div style={{ 
          fontWeight: typography.weights.semibold, 
          color: colors.brand.primary,
          marginBottom: '2px',
          fontSize: typography.sizes.base,
          fontFamily: typography.fonts.primary
        }}>
          {rowData.id}
        </div>
        <div style={{ 
          fontSize: typography.sizes.xs, 
          color: colors.text.secondary,
          fontFamily: typography.fonts.monospace,
          letterSpacing: typography.letterSpacing.wide
        }}>
          {rowData.customId}
        </div>
      </div>
    );
  };

  return (
    <Card 
      style={{ 
        margin: '20px 0',
        border: 'none',
        boxShadow: '0 2px 8px rgba(0,0,0,0.08)',
        borderRadius: '12px'
      }}
    >
      <DataTable 
        value={data} 
        loading={loading}
        paginator 
        rows={rows}
        totalRecords={totalRecords}
        first={first}
        onPage={onPage}
        lazy
        rowsPerPageOptions={[10, 25, 50, 100]}
        tableStyle={{ minWidth: '50rem' }}
        stripedRows={false}
        showGridlines={false}
        emptyMessage="No loads found."
        className="p-datatable-minimalist"
        style={{
          fontSize: '14px',
        }}
        rowClassName={() => 'custom-row'}
      >
        <Column 
          field="id" 
          header="Load ID" 
          body={idTemplate}
          style={{ 
            width: '14rem',
            padding: '1.5rem 1rem',
            borderBottom: `1px solid ${colors.border.light}`
          }}
          headerStyle={{
            backgroundColor: colors.background.secondary,
            color: colors.text.secondary,
            fontWeight: typography.weights.medium,
            fontSize: typography.sizes.xs,
            textTransform: 'uppercase' as const,
            letterSpacing: typography.letterSpacing.wide,
            padding: '1.25rem 1rem',
            borderBottom: `2px solid ${colors.border.light}`,
            fontFamily: typography.fonts.primary
          }}
        />
        
        <Column 
          field="status.code.value" 
          header="Status" 
          body={statusTemplate}
          style={{ 
            width: '10rem',
            padding: '1.5rem 1rem',
            borderBottom: '1px solid #f1f3f4'
          }}
          headerStyle={{
            backgroundColor: '#fafbfc',
            color: '#5f6368',
            fontWeight: '500',
            fontSize: '13px',
            textTransform: 'uppercase',
            letterSpacing: '0.5px',
            padding: '1.25rem 1rem',
            borderBottom: '2px solid #e8eaed'
          }}
        />
        
        <Column 
          field="customerOrder.0.customer.name" 
          header="Customer" 
          body={customerTemplate}
          style={{ 
            width: '18rem',
            padding: '1.5rem 1rem',
            borderBottom: '1px solid #f1f3f4'
          }}
          headerStyle={{
            backgroundColor: '#fafbfc',
            color: '#5f6368',
            fontWeight: '500',
            fontSize: '13px',
            textTransform: 'uppercase',
            letterSpacing: '0.5px',
            padding: '1.25rem 1rem',
            borderBottom: '2px solid #e8eaed'
          }}
        />
        
        <Column 
          field="created" 
          header="Created" 
          body={dateTemplate}
          style={{ 
            width: '12rem',
            padding: '1.5rem 1rem',
            borderBottom: '1px solid #f1f3f4'
          }}
          headerStyle={{
            backgroundColor: '#fafbfc',
            color: '#5f6368',
            fontWeight: '500',
            fontSize: '13px',
            textTransform: 'uppercase',
            letterSpacing: '0.5px',
            padding: '1.25rem 1rem',
            borderBottom: '2px solid #e8eaed'
          }}
        />
        
        <Column 
          field="lastUpdatedOn" 
          header="Last Updated" 
          body={updatedTemplate}
          style={{ 
            width: '12rem',
            padding: '1.5rem 1rem',
            borderBottom: '1px solid #f1f3f4'
          }}
          headerStyle={{
            backgroundColor: '#fafbfc',
            color: '#5f6368',
            fontWeight: '500',
            fontSize: '13px',
            textTransform: 'uppercase',
            letterSpacing: '0.5px',
            padding: '1.25rem 1rem',
            borderBottom: '2px solid #e8eaed'
          }}
        />
        
        <Column 
          field="carrierOrder" 
          header="Carrier" 
          body={(rowData) => rowData.carrierOrder ? 'Assigned' : 'Unassigned'}
          style={{ 
            width: '10rem',
            padding: '1.5rem 1rem',
            borderBottom: '1px solid #f1f3f4'
          }}
          headerStyle={{
            backgroundColor: '#fafbfc',
            color: '#5f6368',
            fontWeight: '500',
            fontSize: '13px',
            textTransform: 'uppercase',
            letterSpacing: '0.5px',
            padding: '1.25rem 1rem',
            borderBottom: '2px solid #e8eaed'
          }}
        />
      </DataTable>
    </Card>
  );
};

export default TableView;
