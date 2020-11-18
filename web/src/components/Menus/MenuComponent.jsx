import React from 'react';
import MenuItem from "../UI/MenuItem/MenuItem";
import PropTypes from 'prop-types';

const MenuComponent = (props) => {
    return (
        <React.Fragment>
            <MenuItem linkTo={'/cart'}>
                Cart <span className="badge badge-light">{props.cartCount}</span>
            </MenuItem>
        </React.Fragment>
    )
};

MenuComponent.propTypes = {
    cartCount: PropTypes.number
};

export default MenuComponent;
