import React, { useState } from 'react';
import { CContainer, CRow, CCol, CButton, CForm, CFormLabel, CFormInput } from '@coreui/react';

export default function HomePageContainer() {
    const [value, setValue] = useState('');

    const _onChange = event => {
        setValue(event.target.value);
    };

    const _onSubmit = event => {
        event.preventDefault();
        event.stopPropagation();
        alert('call shorten url api');
    };

    return (
        <CContainer fluid>
            <CRow className="my-4">
                <p className="fs-4 text-center">Enter the link you want to shorten</p>
            </CRow>
            <CRow>
                <CForm className="row g-3" onSubmit={_onSubmit}>
                    <CRow className="justify-content-center">
                        <CCol md={8}>
                            <CFormLabel htmlFor="shortUrlInput" className="visually-hidden">
                                Link
                            </CFormLabel>
                            <CFormInput
                                required
                                id="shortUrlInput"
                                placeholder="Enter link here"
                                feedbackInvalid="Please provide a valid link"
                                value={value}
                                onChange={_onChange}
                            />
                        </CCol>
                        <CCol xs="auto">
                            <CButton disabled={!value} type="submit" color="primary">
                                Shorten url
                            </CButton>
                        </CCol>
                    </CRow>
                </CForm>
            </CRow>
        </CContainer>
    );
}
