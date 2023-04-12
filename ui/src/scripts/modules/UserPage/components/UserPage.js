import React from 'react';
import { CContainer, CRow, CCol, CTable } from '@coreui/react';

const columns = [
    {
        key: 'longUrl',
        label: 'Original url',
        _props: { scope: 'col' },
    },
    {
        key: 'shortUrl',
        label: 'Short url',
        _props: { scope: 'col' },
    },
];

const UserPage = ({ urls = [] }) => {
    const items = urls.map((url = {}) => ({
        longUrl: (
            <a href={url.redirectUrl} target="_blank" rel="noreferrer">
                {url.redirectUrl}
            </a>
        ),
        shortUrl: (
            <a href={url.encodedUrl} target="_blank" rel="noreferrer">
                {url.encodedUrl}
            </a>
        ),
        _cellProps: { id: { scope: 'row' } },
    }));

    console.log('items', items);
    return (
        <CContainer fluid>
            <CRow className="my-4">
                <h2 className="text-center">My urls</h2>
            </CRow>
            <CRow className="justify-content-center">
                <CCol md={8}>
                    <CTable columns={columns} items={items} />
                </CCol>
            </CRow>
        </CContainer>
    );
};

export default UserPage;
