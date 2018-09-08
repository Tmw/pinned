import { inject, observer } from "mobx-react";
import React from "react";
import { Views } from "Constants";

import StarBox from "components/StarBox";
import PageTitle from "components/PageTitle";

class ScoreView extends React.Component {
  get challengeStore() {
    return this.props.store.ChallengeStore;
  }

  get viewStore() {
    return this.props.store.ViewStore;
  }

  handleRetryClicked = () => {
    this.challengeStore.fetchChallenges();
    this.viewStore.presentView(Views.LOADING);
  };

  render() {
    const score = this.challengeStore.score;
    const numChallenges = this.challengeStore.challenges.length;

    return (
      <React.Fragment>
        <PageTitle subtitle="Your report card!" />

        <div className="score-view">
          <div className="score--box">
            <StarBox stars={score} total={numChallenges} />

            <p>
              You scored {score} out of {numChallenges}!
            </p>

            <div className="center">
              <button onClick={this.handleRetryClicked}>Try again</button>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(ScoreView));
