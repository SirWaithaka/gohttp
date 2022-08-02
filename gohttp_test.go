package gohttp

import "testing"

func TestStatusIsSuccess(t *testing.T) {

	t.Run("test that it returns true if status is 2**", func(t *testing.T) {
		r := StatusIsSuccess(202)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsSuccess(200)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsSuccess(299)
		if !r {
			t.Errorf("should return true")
		}

	})

	t.Run("test that it returns false if status is 3** or 4** or 5**", func(t *testing.T) {
		r := StatusIsSuccess(300)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(400)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(404)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(500)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsSuccess(503)
		if r {
			t.Errorf("should return false")
		}

	})
}

func TestStatusIsClientError(t *testing.T) {

	t.Run("test that it returns true if status is 4**", func(t *testing.T) {
		r := StatusIsClientError(400)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsClientError(404)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsClientError(455)
		if !r {
			t.Errorf("should return true")
		}

	})

	t.Run("test that it returns false if status is 2** or 4** or 5**", func(t *testing.T) {
		r := StatusIsClientError(300)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(200)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(204)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(500)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsClientError(503)
		if r {
			t.Errorf("should return false")
		}

	})
}

func TestStatusIsServerError(t *testing.T) {

	t.Run("test that it returns true if status is 5**", func(t *testing.T) {
		r := StatusIsServerError(500)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsServerError(503)
		if !r {
			t.Errorf("should return true")
		}

		r = StatusIsServerError(599)
		if !r {
			t.Errorf("should return true")
		}

	})

	t.Run("test that it returns false if status is 2** or 4** or 3**", func(t *testing.T) {
		r := StatusIsServerError(300)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(311)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(200)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(204)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(400)
		if r {
			t.Errorf("should return false")
		}

		r = StatusIsServerError(403)
		if r {
			t.Errorf("should return false")
		}

	})
}
