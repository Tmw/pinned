import { types, getRoot } from "mobx-state-tree";
import Author from "state/Author";
import { Views } from "Constants";

const Challenge = types
  .model({
    id: types.string,
    text: types.string,
    choices: types.array(Author),
    author: Author,
    answeredAuthorId: types.maybeNull(types.string)
  })
  .actions(self => ({
    answer(authorId) {
      const { ChallengeStore, ViewStore } = getRoot(self);

      self.answeredAuthorId = authorId;

      if (!ChallengeStore.currentChallenge) {
        ViewStore.presentView(Views.SCORE);
      }
    }
  }))
  .views(self => ({
    get isAnswered() {
      return self.answeredAuthorId !== null;
    },

    get isCorrect() {
      return self.answeredAuthorId === self.author.id;
    }
  }));

export default Challenge;
