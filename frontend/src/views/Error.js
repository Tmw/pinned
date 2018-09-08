import React from "react";
import { inject, observer } from "mobx-react";
import { Helmet } from "react-helmet";

class ErrorView extends React.Component {
  render() {
    const { error } = this.props.store.ViewStore;

    return (
      <React.Fragment>
        <Helmet title="Pinned! - Uh-Oh!" />

        <div className="error-view">
          <h1>Something went wrong! :(</h1>
          <span className="error--details">{error}</span>
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(ErrorView));
