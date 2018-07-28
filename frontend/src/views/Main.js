import React, { Component } from "react";
import { inject, observer } from "mobx-react";

import { Views } from "../Constants";
import LoaderView from "./Loader";
import ErrorView from "./Error";
import PinView from "./Pin";
import ScoreView from "./Score";

class Main extends Component {
  componentDidMount() {
    this.props.store.fetchChallanges();
  }

  renderChildView() {
    switch (this.props.store.currentView) {
      default:
      case Views.LOADING:
        return <LoaderView />;

      case Views.ERROR:
        return <ErrorView />;

      case Views.PIN:
        return <PinView />;

      case Views.SCORE:
        return <ScoreView />;
    }
  }

  onNextView = e => {
    this.props.store.showNextView();
  };

  render() {
    return (
      <div className="App">
        {this.renderChildView()}

        <br />
        <br />
        <br />
        <button onClick={this.onNextView}>Next State</button>
      </div>
    );
  }
}

export default inject("store")(observer(Main));
