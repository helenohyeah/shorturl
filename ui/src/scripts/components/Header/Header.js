import React, { useState } from 'react';
import {
    CNavbar,
    CNavbarToggler,
    CNavbarBrand,
    CNavbarText,
    CCollapse,
    CContainer,
    CNavbarNav,
    CNavItem,
    CNavLink,
} from '@coreui/react';
import { useHistory } from 'react-router-dom';

const Header = ({ userId }) => {
    const history = useHistory();
    const [visible, setVisible] = useState(false);
    const isLoggedIn = !!userId;

    const navItems = isLoggedIn
        ? []
        : [
              {
                  path: '/login',
                  label: 'Login',
              },
              {
                  path: '/register',
                  label: 'Register',
              },
              // for ease of testing
              {
                  path: '/u/123',
                  label: 'User 123',
              },
              {
                  path: '/not_found',
                  label: 'Not found',
              },
              {
                  path: '/anything',
                  label: 'Anything',
              },
          ];

    const _onClick = path => event => {
        event.preventDefault();
        history.push(path);
    };

    return (
        <CNavbar expand="md" colorScheme="dark" className="bg-primary">
            <CContainer fluid>
                <CNavbarBrand href="/" alt="Short url home page" onClick={_onClick('/')}>
                    ShortURL
                </CNavbarBrand>
                {isLoggedIn && <CNavbarText>{`Logged in as ${userId}`}</CNavbarText>}
                {navItems.length && (
                    <>
                        <CNavbarToggler
                            aria-label="Toggle navigation"
                            aria-expanded={visible}
                            onClick={() => setVisible(!visible)}
                        />

                        <CCollapse className="navbar-collapse" visible={visible}>
                            <CNavbarNav>
                                {navItems.map((item = {}) => (
                                    <CNavItem key={item.label}>
                                        <CNavLink href={item.path} onClick={_onClick(item.path)}>
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
