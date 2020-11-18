
import React, { Component } from "react";
import axios from "axios";
import HomeSale from "./components/HomeSale";
import Loader from "../../components/Loader/Index";
import "./Home.css";

class Home extends Component {
  constructor() {
    super();
    this.state = {
      loading: true,
      productsProps: []
    };
  }

  componentDidMount() {
    
      axios.get(`http://localhost:8081/products`)
        .then(res => {
          const productsProps = res.data;
          this.setState({
            loading: false,
            productsProps
          });
        }).then(err => {
          this.setState({
            error: err,
            loading: false,
          })
        }).finally(() => {
          this.setState({
            loading: false,
          })
        })
  }

  render() {
    return this.state.loading ? (
      <Loader />
    ) : (
        <React.Fragment>
          <HomeSale product={this.state.product} />
        </React.Fragment>
      );
  }
}

export default Home;
