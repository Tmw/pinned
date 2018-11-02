import React from "react";
import PageTitle from "components/PageTitle";
import "stylesheets/empty-view.css";

class EmptyView extends React.Component {
  render() {
    return (
      <React.Fragment>
        <PageTitle subtitle="No pins yet!" />

        <div className="view empty-view">
          <h1>No pins yet..</h1>
          <span className="empty--details">
            Start pinning some items or select an alternate channel.
          </span>
        </div>
      </React.Fragment>
    );
  }
}

export default EmptyView;
