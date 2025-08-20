import React from 'react';
import { Card } from 'primereact/card';

const Body: React.FC = () => {
    return (
        <div style={{ display: 'flex', height: '100vh' }}>
            {/* Left Tab */}
            <div style={{
                width: '400px',
                background: '#222',
                color: '#fff',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'
            }}>
                <span>Tab</span>
            </div>
            {/* Main Body */}
            <div style={{
                flex: 1,
                background: '#f5f5f5',
                padding: '24px'
            }}>
                <Card 
                    title="Drumkit Info" 
                    subTitle="PrimeReact Card Example"
                    style={{
                        background: '#fff',
                        boxShadow: '0 4px 16px rgba(0,0,0,0.12)',
                        borderRadius: '8px',
                        padding: '20px',
                    }}
                >
                    <p>
                        This is a simple drumkit application. 
                        <br />
                        <strong>Features:</strong>
                        <ul>
                            <li>Play drum sounds</li>
                            <li>Customizable pads</li>
                            <li>Responsive layout</li>
                        </ul>
                    </p>
                </Card>
                <Card 
                    title="Drumkit Info" 
                    subTitle="PrimeReact Card Example"
                    style={{
                        background: '#fff',
                        boxShadow: '0 4px 16px rgba(0,0,0,0.12)',
                        borderRadius: '8px',
                        padding: '20px',
                    }}
                >
                    <p>
                        This is a simple drumkit application. 
                        <br />
                        <strong>Features:</strong>
                        <ul>
                            <li>Play drum sounds</li>
                            <li>Customizable pads</li>
                            <li>Responsive layout</li>
                        </ul>
                    </p>
                </Card>
            </div>
        </div>
    );
};

export default Body;