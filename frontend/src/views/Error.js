import React from "react";
import { inject, observer } from "mobx-react";

class ErrorView extends React.Component {
  render() {
    const { error } = this.props.store.ViewStore;

    return (
      <div className="error-view">
        <h1>Something went wrong! :(</h1>
        <span className="error--details">{error}</span>
      </div>
    );
  }
}

export default inject("store")(observer(ErrorView));
