<template>
  <ticket-details :is-edit="false" />
</template>

<script>
import TicketDetails from './components/TicketDetails'

export default {
  name: 'TicketAdd',
  components: { TicketDetails }
}
</script>

