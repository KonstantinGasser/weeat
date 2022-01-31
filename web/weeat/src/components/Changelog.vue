<template>
  <div class="content">
    <div class="headline text-center">
      <h1>
        <a href="/"> <span>we</span>Eat </a>
      </h1>
    </div>
    <section>
        <div class="space"></div>
        <h2 class="text-center"><span class="inline-name"><span>we</span>Eat</span> - Changelog 🎉</h2>
        <div class="info-card" v-for="release in releases" :key="release?.id" @click="setActive(release?.id)">
            <h3>Release {{release.version}} <small>{{release.date}}</small></h3>
            <small>
                <i v-if="selectedRelease !== release.id" class="bi bi-caret-right-fill">show more</i>
                <i v-if="selectedRelease === release.id" class="bi bi-caret-down-fill">show less</i>
            </small>
            <div class="changes" :class="{'expand': selectedRelease === release.id}">
                <div class="change-card" v-for="change, index in release?.changes" :key="change.id">
                    <span class="change-title text-emph">👉 {{change.title}}</span>
                    <p class="change-desc">{{change.description}}</p>
                <hr v-if="index < release?.changes?.length-1">
                </div>
            </div>
        </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "Changelog",
  data() {
    return {
      selectedRelease: "",
      releases: [
        {
          id: "cf5d6h",
          date: "2022-01-31",
          version: "cf5d6h",
          changes: [
            {
              id: 0,
              title: "Adding info pages to the website",
              description:
                "adding pages to q&a where one can look up how the app works and see about privacy information. adding pages for change-log content :D",
            },
            {
              id: 1,
              title: "Adding info pages to the website",
              description:
                "adding pages to q&a where one can look up how the app works and see about privacy information. adding pages for change-log content :D",
            },
          ],
        },
      ],
    };
  },
  methods: {
      setActive(release) {
          if (this.selectedRelease === release) {
              this.selectedRelease = ""
              return 
          }
          this.selectedRelease = release
      }
  },
};
</script>

<style scoped>
.content {
  padding: 15px 0px;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

section {
  width: 100vw;
  height: 100%;
  overflow: scroll;

  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 15px 10px;
}

section .space {
    margin: 10px 0;
}

section .info-card {
  box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
  background: var(--box-bg);
  border-radius: 14px;
  padding: 15px;
}

section .info-card h3 {
    display: flex;
    justify-content: space-between;
    align-items: flex-end;
}

section .info-card .changes {
  opacity: 0;
  height: 0px;
  visibility: hidden;
  transition: height 100ms linear;
  transition: opacity 50ms linear;
}

section .info-card > .changes.expand {
  height: max-content;
  visibility: visible;
  opacity: 1;
  transition: opacity 300ms linear;
}

</style>
