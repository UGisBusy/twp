import 'bootstrap/dist/css/bootstrap.min.css';
import '@components/style.css';
import '@style/global.css';

import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Row, Col, NavbarBrand, Button, Dropdown } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faCartShopping, faSearch } from '@fortawesome/free-solid-svg-icons';
import { Link, createSearchParams, useNavigate, useSearchParams } from 'react-router-dom';
import { FormEvent, useState } from 'react';

import LogoImgUrl from '@assets/images/logo.png';

import { useAuth } from '@lib/Auth';

const DropDownStyle = {
  borderRadius: '25px',
  border: '1px solid var(--border)',
  background: ' var(--layout)',
  padding: '10% 5% 10% 5%',
  color: 'white',
};

const DropButtonStyle = {
  background: ' var(--layout)',
  border: 'none',
};

type sortByType = 'price' | 'stock' | 'sales' | 'relevancy';
type orderType = 'asc' | 'desc';

const NavBar = () => {
  const token = useAuth();
  const navigate = useNavigate();
  const [searchString, setSearchString] = useState<string>('');
  const [searchParams] = useSearchParams();
  // TODO: read user auth later
  const isAdmin = true;

  const logout = async () => {
    await fetch('/api/oauth/logout', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    navigate('/login');
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const params = {
      q: searchString,
      sortBy: (searchParams.get('sortBy') as sortByType) ?? 'relevancy',
      order: (searchParams.get('order') as orderType) ?? 'desc',
      minPrice: searchParams.get('minPrice') ?? (0).toString(),
      maxPrice: searchParams.get('maxPrice') ?? (100000).toString(),
      minStock: searchParams.get('minStock') ?? (0).toString(),
      maxStock: searchParams.get('maxStock') ?? (100000).toString(),
      haveCoupon: searchParams.get('haveCoupon') ?? false.toString(),
    };
    navigate({ pathname: '/search', search: `?${createSearchParams(params)}` });
  };
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchString(e.target.value);
  };

  return (
    <div className='navbar_twp'>
      <form onSubmit={handleSubmit}>
        <Navbar expand='xl' style={{ padding: '0px 8% 0px 8%' }}>
          <Row style={{ width: '100%' }}>
            <Col xs={2} className='center'>
              <NavbarBrand href='/' className='disappear_desktop'>
                <img src={LogoImgUrl} alt='logo' style={{ width: '35px' }} />
              </NavbarBrand>
            </Col>
            <Col xs={8} className='center'>
              <Nav style={{ width: '100%' }}>
                <div className='disappear_desktop'>
                  <div className='input_container'>
                    <input
                      type='text'
                      placeholder='Search'
                      className='search'
                      value={searchString}
                      onChange={handleChange}
                    />
                    <FontAwesomeIcon icon={faSearch} className='search_icon' />
                  </div>
                </div>
              </Nav>
            </Col>

            <Col xs={2} ld={12} className='center'>
              <Navbar.Toggle aria-controls='seller-nav' />
            </Col>
            <Col xs={12} md={12}>
              <Navbar.Collapse id='seller-nav' className='navbar-dark'>
                <Row style={{ width: '100%' }}>
                  <Col xs={4}>
                    <Nav className='mt-auto'>
                      <Dropdown>
                        <Dropdown.Toggle
                          id='dropdown-custom-1'
                          style={DropButtonStyle}
                          className='nav_link'
                        >
                          Sell
                        </Dropdown.Toggle>
                        <Dropdown.Menu style={DropDownStyle}>
                          <Link
                            to='/user/seller/info'
                            className='none nav_link'
                            style={{ padding: '0' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>Shop Info</div>
                          </Link>
                          <Link
                            to='/user/seller/manageProducts'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>All Products</div>
                          </Link>
                          <Link
                            to='/user/seller/manageCoupons'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>All Coupons</div>
                          </Link>
                          <Link
                            to='/user/seller/orders'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>All Shipments</div>
                          </Link>
                          <Link
                            to='/user/seller/reports'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>All Reports</div>
                          </Link>
                        </Dropdown.Menu>
                      </Dropdown>
                      <Link to='/coupons' className='nav_link none' style={{ paddingLeft: '10px' }}>
                        Coupons
                      </Link>
                    </Nav>
                  </Col>
                  <Col xs={4} />
                  <Col xs={4} className='right'>
                    <Nav className='ms-auto'>
                      <Dropdown>
                        <Dropdown.Toggle
                          id='dropdown-custom-1'
                          style={DropButtonStyle}
                          className='nav_link'
                        >
                          <FontAwesomeIcon icon={faUser} />
                        </Dropdown.Toggle>
                        <Dropdown.Menu style={DropDownStyle}>
                          <Link to='/user/info' className='none nav_link' style={{ padding: '0' }}>
                            <div style={{ padding: '5px 10% 5px 10%' }}>Personal Info</div>
                          </Link>
                          <Link
                            to='/user/security'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>Security</div>
                          </Link>
                          <Link
                            to='/user/buyer/order'
                            className='none nav_link'
                            style={{ padding: '0%' }}
                          >
                            <div style={{ padding: '5px 10% 5px 10%' }}>Order History</div>
                          </Link>
                          <hr
                            style={{
                              padding: '0',
                              margin: '5px',
                              color: 'var(--border)',
                              opacity: '1',
                            }}
                          />
                          {isAdmin ? (
                            <>
                              <Link to='/admin' className='nav_link none' style={{ padding: '0%' }}>
                                <div style={{ padding: '5px 10% 5px 10%' }}>Admin</div>
                              </Link>
                              <hr
                                style={{
                                  padding: '0',
                                  margin: '5px',
                                  color: 'var(--border)',
                                  opacity: '1',
                                }}
                              />
                            </>
                          ) : (
                            <></>
                          )}
                          <div
                            className='none nav_link'
                            style={{ padding: '5px 10% 5px 10%', cursor: 'pointer' }}
                            onClick={logout}
                          >
                            Logout
                          </div>
                        </Dropdown.Menu>
                      </Dropdown>
                      <Link
                        to='/buyer/cart'
                        className='nav_link none'
                        style={{ paddingLeft: '10px' }}
                      >
                        <FontAwesomeIcon icon={faCartShopping} />
                      </Link>
                    </Nav>
                  </Col>
                </Row>
              </Navbar.Collapse>
            </Col>
          </Row>
        </Navbar>

        <div className='disappear_phone disappear_tablet'>
          <hr style={{ color: 'white', opacity: '0.5', margin: '5px' }} />

          <Row className='center' style={{ padding: '0px 8% 0px 8%' }}>
            <Col sm={3}>
              <Link to='/' className='none'>
                <div className='center_vertical'>
                  <img src={LogoImgUrl} alt='logo' style={{ width: '35px' }} />
                  &nbsp;&nbsp; <span className='nav_title'>Too White Powder</span>
                </div>
              </Link>
            </Col>
            <Col sm={6}>
              <div className='input_container'>
                <input
                  type='text'
                  placeholder='Search'
                  className='search'
                  value={searchString}
                  onChange={handleChange}
                />
                <FontAwesomeIcon icon={faSearch} className='search_icon' />
              </div>
            </Col>
            <Col sm={3}>
              <Button className='search_button center' type='submit'>
                Search
              </Button>
            </Col>
          </Row>
        </div>
      </form>
    </div>
  );
};

export default NavBar;
