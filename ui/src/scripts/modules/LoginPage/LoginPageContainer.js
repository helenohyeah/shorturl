import React from 'react';
import { useSelector } from 'react-redux';
import {
    CContainer,
    CRow,
    CCol,
    CButton,
    CForm,
    CFormInput,
    CFormLabel,
    CInputGroup,
    CInputGroupText,
} from '@coreui/react';
import { useLoginPageActions } from './reducer/loginPage';

const getState = state => state.login;

const LoginPageContainer = ({ isRegister }) => {
    const { onFormChange, onToggleShowPassword } = useLoginPageActions();

    const isPasswordHidden = useSelector(state => getState(state).isPasswordHidden);
    const username = useSelector(state => getState(state).username);
    const password = useSelector(state => getState(state).password);

    const _onFormChange = name => event => {
        const value = event.target.value;
        onFormChange(name, value);
    };

    const _onSubmit = event => {
        event.preventDefault();
        event.stopPropagation();
        alert('call login / register');
    };

    return (
        <CContainer fluid>
            <CRow className="my-4">
                <p class="fs-4 text-center">{isRegister ? 'Register' : 'Login'}</p>
            </CRow>
            <CRow>
                <CForm className="row g-3" onSubmit={_onSubmit}>
                    <CRow className="justify-content-center">
                        <CCol md={8}>
                            <CFormLabel htmlFor="usernameInput">Username</CFormLabel>
                            <CFormInput
                                required
                                id="usernameInput"
                                placeholder="Enter username"
                                feedbackInvalid="Please provide username"
                                value={username}
                                onChange={_onFormChange('username')}
                            />
                            <CFormLabel className="mt-2" htmlFor="passwordInput">
                                Password
                            </CFormLabel>
                            <CInputGroup>
                                <CFormInput
                                    required
                                    type={isPasswordHidden ? 'password' : 'text'}
                                    id="passwordInput"
                                    placeholder="Enter password"
                                    aria-describedby="passwordInput"
                                    feedbackInvalid="Please provide a password"
                                    value={password}
                                    onChange={_onFormChange('password')}
                                />
                                <CInputGroupText onClick={onToggleShowPassword}>
                                    {isPasswordHidden ? 'Show' : 'Hide'}
                                </CInputGroupText>
                            </CInputGroup>
                            <div className="d-grid gap-2">
                                <CButton
                                    className="mt-4"
                                    disabled={!username && !password}
                                    type="submit"
                                    color="primary"
                                >
                                    {isRegister ? 'Register' : 'Login'}
                                </CButton>
                            </div>
                        </CCol>
                    </CRow>
                </CForm>
            </CRow>
        </CContainer>
    );
};

export default LoginPageContainer;
