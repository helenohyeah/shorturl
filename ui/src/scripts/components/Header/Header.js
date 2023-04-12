import React, { useState } from 'react';
import {
    CNavbar,
    CNavbarToggler,
    CNavbarBrand,
    CCollapse,
    CContainer,
    CNavbarNav,
    CNavItem,
    CNavLink,
} from '@coreui/react';
import { useHistory } from 'react-router-dom';

const Header = ({ userId, items = [] }) => {
    const history = useHistory();
    const [visible, setVisible] = useState(false);

    const _onClick = item => event => {
        event.preventDefault();
        if (item.onClick) {
            item.onClick();
        }
        history.push(item.path);
    };

    return (
        <CNavbar expand="md" colorScheme="dark" className="bg-primary">
            <CContainer fluid>
                <CNavbarBrand href="/" alt="Short url home page" onClick={_onClick({ path: '/' })}>
                    ShortURL
                </CNavbarBrand>
                {items.length && (
                    <>
                        <CNavbarToggler
                            aria-label="Toggle navigation"
                            aria-expanded={visible}
                            onClick={() => setVisible(!visible)}
                        />

                        <CCollapse className="navbar-collapse" visible={visible}>
                            <CNavbarNav className="me-auto">
                                {items.map((item = {}) => (
                                    <CNavItem key={item.label}>
                                        <CNavLink href={item.path} onClick={_onClick(item)}>
                                            {item.label}
                                        </CNavLink>
                                    </CNavItem>
                                ))}
                            </CNavbarNav>
                        </CCollapse>
                    </>
                )}
            </CContainer>
        </CNavbar>
    );
};

export default Header;
