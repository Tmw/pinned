import React, { Component } from "react";
import { inject, observer } from "mobx-react";

import { Views } from "Constants";
import LoaderView from "views/Loader";
import ErrorView from "views/Error";
import PinView from "views/Pin";
import ScoreView from "views/Score";

class Main extends Component {
  componentDidMount() {
    this.props.store.ChallengeStore.fetchChallenges();
  }

  renderChildView() {
    const { ViewStore } = this.props.store;

    switch (ViewStore.currentView) {
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

  render() {
    return <div className="App">{this.renderChildView()}</div>;
  }
}

export default inject("store")(observer(Main));
