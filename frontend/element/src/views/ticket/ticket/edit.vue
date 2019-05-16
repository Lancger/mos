<template>
  <ticket-details :is-edit="true" />
</template>

<script>
import TicketDetails from './components/TicketDetails'

export default {
  name: 'TicketEdit',
  components: { TicketDetails }
}
</script>
