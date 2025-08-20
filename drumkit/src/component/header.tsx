import React from 'react';
import { Button } from 'primereact/button';

const Header: React.FC = () => {
    return (

        <header style={{ padding: '1rem', background: '#222', color: '#fff' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', minHeight: '100px' }}>
            <div style={{ flex: '1 1 20%', textAlign: 'left' }}>
                <h1 style={{ margin: 0, fontSize: '2.5rem', display: 'flex', alignItems: 'center' }}>
                🥁 DRUMKIT
                </h1>
            </div>
            <div style={{ flex: '1 1 50%', textAlign: 'left' }}>
                <a href="/" style={{ marginRight: '1rem', color: '#fff', textDecoration: 'none' }}>Home</a>
                <a href="/products" style={{ marginRight: '1rem', color: '#fff', textDecoration: 'none' }}>Products</a>
                <a href="/support" style={{ color: '#fff', textDecoration: 'none' }}>Support</a>
            </div>
            <div style={{ flex: '1 1 0%', textAlign: 'right' }}>
                <Button label="🛒" className="p-button-rounded p-button-text" aria-label="Cart" style={{ marginRight: '0.5rem' }} />
                <Button label="🔍" className="p-button-rounded p-button-text" aria-label="Search" style={{ marginRight: '0.5rem' }} />
                <Button label="🔔" className="p-button-rounded p-button-text" aria-label="Notifications" />
            </div>
            </div>
        </header>
    );
};

export default Header;