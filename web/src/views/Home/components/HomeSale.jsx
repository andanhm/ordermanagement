import React from "react";
import { connect } from "react-redux";
import { NavLink } from "react-router-dom";
import { currencyToUse,
  productPrice,
  productDiscountPrice, } from "../../../Utility/currency";
import { VISIBILITY_FILTERS } from "../../../static/constants";
import { getProductsByFilter } from "../../../store/selectors";

const HomeSale = (props) => {
  let currencyKeys = currencyToUse(props.usedCurrencyProp);

  let products = props.productsProps.map((product, index) => {
    return (
      <div className="card card-body shadow" key={index}>
        <img
          className="card-img-top"
          src={require(`../../../assets/images/shop_images/${product.img}`)}
          alt="product"
        />
        <h5 className="card-title">{product.name}</h5>
        <p className="card-text">
          <div className="product-price-container">
            <span className="product-price">
              {currencyKeys.name}
              {productPrice(
                product.price,
                currencyKeys.value
              )}
            </span>
            {product.discount_price ? (
              <span className="product-discount-price">
                {currencyKeys.name}
                {productPrice(
                  product.discount_price,
                  currencyKeys.value
                )}
              </span>
            ) : null}
            {product.discount_price ? (
              <span className="product-percentage-discount">
                {productDiscountPrice(
                  product.price,
                  product.discount_price
                )}
              </span>
            ) : null}
          </div>
        </p>
        <NavLink
          className="btn btn-primary btn-sm"
          to={`/product/${product.slug}`}
          exact
        >
          View Item
        </NavLink>
      </div>
    );
  });
  return (
    <div className="container products-section mb-4">
      <div className="products-section-title pb-3">
      </div>
      <div className="products-container">{products}</div>
    </div>
  );
};

const mapStateToProps = (state) => {
  return {
    productsProps: getProductsByFilter(state, VISIBILITY_FILTERS.SALE, 6),
    usedCurrencyProp: state.usedCurrency,
  };
};

export default connect(mapStateToProps)(HomeSale);
