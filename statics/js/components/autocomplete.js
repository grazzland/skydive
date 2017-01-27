/* jshint multistr: true */

Vue.component('autocomplete', {

  template: '\
    <div :class="{\'open\':openSuggestion}">\
        <input class="form-control input-sm" type="text"\
               :placeholder="placeholder"\
               :value="value"\
               @keydown.enter="enter($event.target.value)"\
               @keydown.tab="complete"\
               @keydown.down="down"\
               @keydown.up="up"\
               @input="change($event.target.value)"\
        />\
        <ul class="dropdown-menu">\
            <li v-for="(suggestion, index) in matches"\
                :class="{\'active\': isActive(index)}"\
                @click="click(index)"\
            >\
                <a href="#">{{ suggestion }}</a>\
            </li>\
        </ul>\
    </div>',

  props: {

    value: {
      type: String,
      required: true,
    },

    // Function must return
    // a promise
    suggestions: {
      type: Function,
      required: true,
    },

    placeholder: {
      type: String,
    },

  },

  data: function() {
    return {
      open: false,
      current: 0,
      fetchedSuggestions: [],
    };
  },

  computed: {

    openSuggestion: function() {
      return this.value !== "" &&
             this.matches.length !== 0 &&
             this.open === true;
    },

    matches: function() {
      var self = this;
      return this.fetchedSuggestions.filter(function(s) {
        return s.indexOf(self.value) >= 0;
      });
    },

  },

  methods: {

    fetchSuggestions: function() {
      var self = this;
      this.suggestions()
        .then(function(data) {
          self.fetchedSuggestions = data;
        });
    },

    enter: function(value) {
      this.complete();
      this.$emit('select', value);
    },

    complete: function() {
      if (this.openSuggestion === true) {
        value = this.matches[this.current] || this.value;
        this.open = false;
        this.$emit('input', value);
      }
    },

    click: function(index) {
      this.open = false;
      this.$emit('input', this.matches[index]);
      this.$emit('select', this.matches[index]);
    },

    up: function() {
      if (this.current > 0)
        this.current--;
    },

    down: function() {
      if (this.current < this.matches.length - 1)
        this.current++;
    },

    isActive: function(index) {
      return index == this.current;
    },

    change: function(value) {
      // FIXME: debounce
      this.fetchSuggestions();
      if (this.open === false) {
        this.open = true;
        this.current = 0;
      }
      this.$emit('input', value);
    }

  }

});
