import React, { useState } from 'react';
import { useViewLoadsQuery, useCreateLoadMutation } from '../features/drumkitAPI';
import TableView from './tableView';
import NavigationSidebar from './NavigationSidebar';
import ErrorDisplay from './ErrorDisplay';
import ActionBar from './ActionBar';
import CreateLoadDialog from './CreateLoadDialog';
import { colors } from '../styles/typography';

const Body: React.FC = () => {
    const [showDialog, setShowDialog] = useState(false);
    const [first, setFirst] = useState(0);
    const [rows, setRows] = useState(25);

    // Make the query call with pagination parameters
    const { data, error, isLoading } = useViewLoadsQuery({
        start: first.toString(),
        pageSize: rows.toString()
    });

    const [createLoad, { isLoading: isCreatingLoad }] = useCreateLoadMutation();

    const handlePageChange = (event: any) => {
        setFirst(event.first);
        setRows(event.rows);
    };

    const handleCreateLoad = async (loadData: any) => {
        try {
            await createLoad(loadData).unwrap();
            
            // Hide the dialog after successful creation
            setShowDialog(false);
        } catch (error) {
            console.error('Failed to create load:', error);
            throw error;
        }
    };

    return ( 
        <div style={{ display: 'flex', height: '100vh' }}>
            {/* Main Body */}
            <div style={{
                flex: 1,
                background: colors.background.secondary,
                padding: '24px',
                overflow: 'auto'
            }}>
                <ErrorDisplay error={error} />
                
                <ActionBar onCreateLoad={() => setShowDialog(true)} />
                
                {/* Display the table with API data */}
                <TableView 
                    data={data || []} 
                    loading={isLoading}
                    totalRecords={1000} // You may want to get this from API response
                    first={first}
                    rows={rows}
                    onPage={handlePageChange}
                />

                <CreateLoadDialog 
                    visible={showDialog} 
                    onHide={() => setShowDialog(false)}
                    onCreateLoad={handleCreateLoad}
                    isLoading={isCreatingLoad}
                />
            </div>
        </div>
    );
};

export default Body;
