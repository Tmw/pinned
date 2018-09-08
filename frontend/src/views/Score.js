import { inject, observer } from "mobx-react";
import { Helmet } from "react-helmet";
import React from "react";
import { Views } from "Constants";

import Starbox from "components/Starbox";

class ScoreView extends React.Component {
  handleRetryClicked = e => {
    const { ChallengeStore, ViewStore } = this.props.store;

    ChallengeStore.fetchChallenges();
    ViewStore.presentView(Views.LOADING);
  };

  render() {
    const { ChallengeStore } = this.props.store;

    const countScore = (score, challenge) =>
      challenge.isCorrect ? score + 1 : score;

    const totalChallenges = ChallengeStore.challenges.length;

    const totalScore = ChallengeStore.challenges.reduce(countScore, 0);

    return (
      <React.Fragment>
        <Helmet title="Pinned! - Your report card!" />

        <div className="score-view">
          <div className="score--box">
            <Starbox stars={totalScore} total={totalChallenges} />
            <p>
              You scored {totalScore} out of {totalChallenges}!
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
