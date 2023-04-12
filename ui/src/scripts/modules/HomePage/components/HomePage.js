import React from 'react';
import {
    CContainer,
    CRow,
    CCol,
    CButton,
    CForm,
    CFormLabel,
    CFormInput,
    CAlert,
} from '@coreui/react';
import ShortenedUrl from 'scripts/components/ShortenedUrl';

const HomePage = ({
    isShortened,
    longUrl,
    shortUrl,
    error,
    onChange,
    onSubmit,
    onShortenAgain,
}) => {
    const title = isShortened ? 'Your shortened url' : 'Enter the link you want to shorten';
    return (
        <CContainer fluid>
            <CRow className="my-4">
                <h2 className="text-center">{title}</h2>
            </CRow>
            <CRow>
                {isShortened ? (
                    <CRow className="justify-content-center">
                        <CCol md={8}>
                            <ShortenedUrl longUrl={longUrl} shortUrl={shortUrl} />
                            <div className="d-flex justify-content-center">
                                <CButton className="my-4" color="primary" onClick={onShortenAgain}>
                                    Shorten another link
                                </CButton>
                            </div>
                        </CCol>
                    </CRow>
                ) : (
                    <CForm className="row g-3" onSubmit={onSubmit}>
                        <CRow className="justify-content-center">
                            {error && (
                                <CCol md={10}>
                                    <CRow className="my-4">
                                        <CAlert color="danger">Error: {error}</CAlert>
                                    </CRow>
                                </CCol>
                            )}
                            <CCol md={8}>
                                <CFormLabel htmlFor="longUrl" className="visually-hidden">
                                    Link
                                </CFormLabel>
                                <CFormInput
                                    required
                                    id="longUrl"
                                    placeholder="Enter link here"
                                    feedbackInvalid="Please provide a valid link"
                                    value={longUrl}
                                    onChange={onChange('longUrl')}
                                />
                            </CCol>
                            <CCol xs="auto">
                                <CButton disabled={!longUrl} type="submit" color="primary">
                                    Shorten url
                                </CButton>
                            </CCol>
                        </CRow>
                    </CForm>
                )}
            </CRow>
        </CContainer>
    );
};

export default HomePage;
