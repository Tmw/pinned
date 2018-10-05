import React from "react";
import { inject, observer } from "mobx-react";
import PageTitle from "components/PageTitle";
import "stylesheets/error-view.css";

class ErrorView extends React.Component {
  render() {
    const { error } = this.props.store.ViewStore;

    return (
      <React.Fragment>
        <PageTitle subtitle="Uh-Oh!" />

        <div className="error-view">
          <h1>Something went wrong! :(</h1>
          <span className="error--details">{error}</span>
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(ErrorView));
